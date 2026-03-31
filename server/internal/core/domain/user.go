package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/Fioneo/taskflow/internal/core/errors"
)

type User struct {
	ID          int
	Version     int
	Fullname    string
	PhoneNumber *string
}

func NewUser(id int, version int, fullname string, phoneNumber *string) User {
	return User{
		ID:          id,
		Version:     version,
		Fullname:    fullname,
		PhoneNumber: phoneNumber,
	}
}

func NewUserUninitialized(fullname string, phoneNumber *string) User {
	return NewUser(
		UninitializedID, UninitializedVersion, fullname, phoneNumber,
	)
}
func (u *User) Validate() error {
	fullNameLength := len([]rune(u.Fullname))
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf("invalid FullName length: %d, %w", fullNameLength, core_errors.ErrInvalidArgument)
	}
	if u.PhoneNumber != nil {
		phoneNumberLength := len([]rune(*u.PhoneNumber))
		if phoneNumberLength < 10 || phoneNumberLength > 15 {
			return fmt.Errorf("invalid PhoneNumber length: %d, %w", phoneNumberLength, core_errors.ErrInvalidArgument)
		}
		re := regexp.MustCompile(`^\+[0-9]+$`)
		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf("invalid PhoneNumber format: %w", core_errors.ErrInvalidArgument)
		}
	}
	return nil
}

type UserPatch struct {
	Fullname    Nulleable[string]
	PhoneNumber Nulleable[string]
}

func (p *UserPatch) Validate() error {
	if p.Fullname.Set && p.Fullname.Value == nil {
		return fmt.Errorf("fullname can't be patched to null: %w", core_errors.ErrInvalidArgument)
	}
	return nil
}
func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate user patch: %w", err)
	}

	tmp := *u

	if patch.Fullname.Set {
		tmp.Fullname = *patch.Fullname.Value
	}

	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = patch.PhoneNumber.Value
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate patched user: %w", err)
	}

	*u = tmp

	return nil
}
func NewUserPatch(fullname Nulleable[string], phoneNumber Nulleable[string]) UserPatch {
	return UserPatch{
		Fullname:    fullname,
		PhoneNumber: phoneNumber,
	}
}
