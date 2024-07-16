package models

import (
	"context"
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
	"github.com/samber/do"
	"github.com/samber/lo"
)

//
// Entities
//

const MaxContentLength = 1000

type Todo struct {
	ID          string     // TODO ID
	Content     string     // TODO 内容
	CreatedAt   time.Time  // TODO 作成日時
	UpdatedAt   time.Time  // TODO 更新日時
	CompletedAt *time.Time // TODO 完了日時
}

func (t *Todo) Validate() error {
	if utf8.RuneCountInString(t.Content) > MaxContentLength {
		return fmt.Errorf("content is too long")
	}

	return nil
}

func (t *Todo) SetContent(content string) {
	t.Content = content
}

func (t *Todo) Complete() {
	t.CompletedAt = lo.ToPtr(time.Now())
}

func (t *Todo) Uncomplete() {
	t.CompletedAt = nil
}

//
// Repositories
//

type TodoRepository interface {
	Find(ctx context.Context, id string) (*Todo, error)
	Create(ctx context.Context, todo *Todo) error
	Update(ctx context.Context, todo *Todo) (*Todo, error)
	Delete(ctx context.Context, todo *Todo) error
}

//
// Factories
//

type TodoFactory interface {
	Create(options *TodoFactoryOptions) *Todo
}

type TodoFactoryOptions struct {
	Content string // TODO 内容
}

type todoFactory struct{}

func NewTodoFactory(i *do.Injector) (TodoFactory, error) {
	return &todoFactory{}, nil
}

func (f *todoFactory) Create(options *TodoFactoryOptions) *Todo {
	now := time.Now()
	return &Todo{
		ID:          uuid.NewString(),
		Content:     options.Content,
		CreatedAt:   now,
		UpdatedAt:   now,
		CompletedAt: nil,
	}
}
