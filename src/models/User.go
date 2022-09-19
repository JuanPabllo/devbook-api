package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare(step string) error {
	if erro := user.validate(step); erro != nil {
		return erro
	}

	user.formatter()

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("required name")
	}

	if user.Nick == "" {
		return errors.New("required nick")
	}

	if user.Email == "" {
		return errors.New("required email")
	}

	if step == "register" && user.Password == "" {
		return errors.New("required password")
	}

	return nil
}

func (user *User) formatter() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
