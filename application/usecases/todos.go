package usecases

import (
	"context"
	"time"

	"github.com/samber/do"
	"github.com/samber/lo"
	"github.com/tadasi/portfolio/domain/models"
)

type TodoInteractor interface {
	FindTodo(ctx context.Context, input *FindTodoInput) (*FindTodoOutput, error)
	CreateTodo(ctx context.Context, input *CreateTodoInput) (*CreateTodoOutput, error)
	UpdateTodo(ctx context.Context, input *UpdateTodoInput) (*UpdateTodoOutput, error)
	DeleteTodo(ctx context.Context, input *DeleteTodoInput) error
	CompleteTodo(ctx context.Context, input *CompleteTodoInput) (*CompleteTodoOutput, error)
	UncompleteTodo(ctx context.Context, input *UncompleteTodoInput) (*UncompleteTodoOutput, error)
}

type FindTodoInput struct {
	TodoID string
}
type FindTodoOutput struct {
	Todo *models.Todo
}

type CreateTodoInput struct {
	Content string
}
type CreateTodoOutput struct {
	TodoID string
}

type UpdateTodoInput struct {
	TodoID  string
	Content string
}
type UpdateTodoOutput struct {
	Todo *models.Todo
}

type DeleteTodoInput struct {
	TodoID string
}

type CompleteTodoInput struct {
	TodoID string
}
type CompleteTodoOutput struct {
	Todo *models.Todo
}

type UncompleteTodoInput struct {
	TodoID string
}
type UncompleteTodoOutput struct {
	Todo *models.Todo
}

type todoInteractor struct {
	todoFactory    models.TodoFactory
	todoRepository models.TodoRepository
}

func NewTodoInteractor(i *do.Injector) (TodoInteractor, error) {
	return &todoInteractor{
		todoFactory:    do.MustInvoke[models.TodoFactory](i),
		todoRepository: do.MustInvoke[models.TodoRepository](i),
	}, nil
}

func (i *todoInteractor) FindTodo(ctx context.Context, input *FindTodoInput) (*FindTodoOutput, error) {
	todo, err := i.todoRepository.Find(
		ctx,
		input.TodoID,
	)
	if err != nil {
		return nil, err
	}

	return &FindTodoOutput{
		Todo: todo,
	}, nil
}

func (i *todoInteractor) CreateTodo(ctx context.Context, input *CreateTodoInput) (*CreateTodoOutput, error) {
	todo := i.todoFactory.Create(&models.TodoFactoryOptions{
		Content: input.Content,
	})

	if err := i.todoRepository.Create(
		ctx,
		todo,
	); err != nil {
		return nil, err
	}

	return &CreateTodoOutput{
		TodoID: todo.ID,
	}, nil
}

func (i *todoInteractor) UpdateTodo(ctx context.Context, input *UpdateTodoInput) (*UpdateTodoOutput, error) {
	todo, err := i.todoRepository.Find(
		ctx,
		input.TodoID,
	)
	if err != nil {
		return nil, err
	}

	todo.Content = input.Content
	updatedTodo, err := i.todoRepository.Update(
		ctx,
		todo,
	)
	if err != nil {
		return nil, err
	}

	return &UpdateTodoOutput{
		Todo: updatedTodo,
	}, nil
}

func (i *todoInteractor) DeleteTodo(ctx context.Context, input *DeleteTodoInput) error {
	todo, err := i.todoRepository.Find(
		ctx,
		input.TodoID,
	)
	if err != nil {
		return err
	}

	if err := i.todoRepository.Delete(
		ctx,
		todo,
	); err != nil {
		return err
	}

	return nil
}

func (i *todoInteractor) CompleteTodo(ctx context.Context, input *CompleteTodoInput) (*CompleteTodoOutput, error) {
	todo, err := i.todoRepository.Find(
		ctx,
		input.TodoID,
	)
	if err != nil {
		return nil, err
	}

	todo.CompletedAt = lo.ToPtr(time.Now())
	updatedTodo, err := i.todoRepository.Update(
		ctx,
		todo,
	)
	if err != nil {
		return nil, err
	}

	return &CompleteTodoOutput{
		Todo: updatedTodo,
	}, nil
}

func (i *todoInteractor) UncompleteTodo(ctx context.Context, input *UncompleteTodoInput) (*UncompleteTodoOutput, error) {
	todo, err := i.todoRepository.Find(
		ctx,
		input.TodoID,
	)
	if err != nil {
		return nil, err
	}

	todo.CompletedAt = nil
	updatedTodo, err := i.todoRepository.Update(
		ctx,
		todo,
	)
	if err != nil {
		return nil, err
	}

	return &UncompleteTodoOutput{
		Todo: updatedTodo,
	}, nil
}
