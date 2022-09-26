package models

import (
	"errors"
	"strings"
	"time"
)

type Posts struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick uint64    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (posts *Posts) Preparar() error {
	if erro := posts.Validate(); erro != nil {
		return erro
	}

	posts.formatar()
	return nil
}

func (posts *Posts) Validate() error {
	if posts.Title == "" {
		return errors.New("o título é obrigatório e não pode estar em branco")
	}

	if posts.Content == "" {
		return errors.New("o conteúdo é obrigatório e não pode estar em branco")
	}

	return nil
}

func (posts *Posts) formatar() {
	posts.Title = strings.TrimSpace(posts.Title)
	posts.Content = strings.TrimSpace(posts.Content)
}
