package usecases

import (
	"context"

	"github.com/samber/do"
	"github.com/tadasi/portfolio/domain/models"
)

type TodoInteractor interface {
	FindTodo(ctx context.Context, input *FindTodoInput) (*FindTodoOutput, error)
	CreateTodo(ctx context.Context, input *CreateTodoInput) (*CreateTodoOutput, error)
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
