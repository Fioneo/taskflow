package domain

import (
	"fmt"
	"time"

	core_errors "github.com/Fioneo/taskflow/internal/core/errors"
)

type Task struct {
	ID          int
	Version     int
	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
	AuthorID    int
}

func NewTask(
	id int,
	version int,
	title string,
	description *string,
	completed bool,
	createdAt time.Time,
	completedAt *time.Time,
	authorID int,
) Task {
	return Task{
		ID:          id,
		Version:     version,
		Title:       title,
		Description: description,
		Completed:   completed,
		CreatedAt:   createdAt,
		CompletedAt: completedAt,
		AuthorID:    authorID,
	}
}

func NewTaskUninitialized(title string, description *string, authorID int) Task {
	return NewTask(
		UninitializedID,
		UninitializedVersion,
		title,
		description,
		false,
		time.Now(),
		nil,
		authorID,
	)
}
func (t *Task) Validate() error {
	titleLength := len([]rune(t.Title))
	if titleLength < 1 || titleLength > 30 {
		return fmt.Errorf("invalid title length: %d, %w", titleLength, core_errors.ErrInvalidArgument)
	}

	if t.Description != nil {
		descriptionLength := len([]rune(*t.Description))

		if descriptionLength < 1 || descriptionLength > 100 {
			return fmt.Errorf("invalid Description length: %d, %w ", descriptionLength, core_errors.ErrInvalidArgument)
		}
	}

	if t.AuthorID < 1 {
		return fmt.Errorf("authorID must be bigger then 0: %w", core_errors.ErrInvalidArgument)
	}

	if t.Completed {
		if t.CompletedAt == nil {
			return fmt.Errorf("'CompletedAt' can't be 'nil' if 'Completed' == true: %w", core_errors.ErrInvalidArgument)
		}

		if t.CompletedAt.Before(t.CreatedAt) {
			return fmt.Errorf("'Completed at' can't be before 'CreatedAt': %w", core_errors.ErrInvalidArgument)
		}
	} else {
		if t.CompletedAt != nil {
			return fmt.Errorf("'CompletedAt' must be 'nil' if Completed == 'false': %w", core_errors.ErrInvalidArgument)
		}
	}
	return nil
}

type TaskPatch struct {
	Title       Nulleable[string]
	Description Nulleable[string]
	Completed   Nulleable[bool]
}

func NewTaskPatch(
	title Nulleable[string],
	description Nulleable[string],
	completed Nulleable[bool],
) TaskPatch {
	return TaskPatch{
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}
func (p *TaskPatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf("'Title'can't be patched to null: %w", core_errors.ErrInvalidArgument)
	}
	if p.Completed.Set && p.Completed.Value == nil {
		return fmt.Errorf("'Completed' can't be patched null: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}
func (t *Task) ApplyPatch(patch TaskPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("Validate task patch: %w", err)
	}

	tmp := *t

	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}
	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}
	if patch.Completed.Set {
		tmp.Completed = *patch.Completed.Value

		if tmp.Completed {
			completedAt := time.Now()
			tmp.CompletedAt = &completedAt
		} else {
			tmp.CompletedAt = nil
		}
	}
	fmt.Println(tmp)
	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("Validate patched task: %w", err)
	}

	*t = tmp

	return nil
}
