package core_http_utils

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	core_errors "github.com/Fioneo/taskflow/internal/core/errors"
)

func GetIntPathValue(r *http.Request, key string) (int, error) {
	pathValue := r.PathValue(key)
	if pathValue == "" {
		return 0, fmt.Errorf("no key='%s' in path values: %w", key, core_errors.ErrInvalidArgument)
	}
	val, err := strconv.Atoi(pathValue)
	if err != nil {
		return 0, fmt.Errorf("path value='%s', by key='%s' not a valid integer: %v: %w", pathValue, key, err, core_errors.ErrInvalidArgument)

	}

	return val, nil
}
func GetDateQueryParam(r *http.Request, key string) (*time.Time, error) {
	layout := "2006-01-02"
	param := r.URL.Query().Get(key)
	if param == "" {
		return nil, nil
	}

	date, err := time.Parse(layout, param)

	if err != nil {
		return nil, fmt.Errorf("param='%s' by key='%s' not a valid date: %v: %w", param, key, err, core_errors.ErrInvalidArgument)
	}

	return &date, nil
}
