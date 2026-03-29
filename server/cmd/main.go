package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/Fioneo/taskflow/internal/core/logger"
	core_postgres_pool "github.com/Fioneo/taskflow/internal/core/repository/pool"
	core_http_middleware "github.com/Fioneo/taskflow/internal/core/transport/http/middleware"
	core_http_server "github.com/Fioneo/taskflow/internal/core/transport/http/server"
	users_repository "github.com/Fioneo/taskflow/internal/features/users/repository"
	users_service "github.com/Fioneo/taskflow/internal/features/users/service"
	users_transport "github.com/Fioneo/taskflow/internal/features/users/transport"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	logger, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("Failed to init application logger: ", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("Initializing postgres pool connection")

	dbpool, err := core_postgres_pool.NewConnectionPool(core_postgres_pool.NewConfigMust(), ctx)
	if err != nil {
		logger.Fatal("failed connection to postgres pool: ", zap.Error(err))
	}

	defer dbpool.Close()

	logger.Debug("Initializing feature", zap.String("feature", "users"))

	usersRepo := users_repository.NewUserRepostory(dbpool)
	usersService := users_service.NewUsersService(usersRepo)
	usersTransportHTTP := users_transport.NewUsersHTTPHandler(usersService)

	logger.Debug("Initializing HTTP server")

	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)

	apiVersionRouter := core_http_server.NewApiVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	httpServer.RegisterApiRoutes(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
