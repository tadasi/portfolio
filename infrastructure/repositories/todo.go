package repositories

import (
	"context"

	"github.com/samber/do"
	"github.com/tadasi/portfolio/domain/models"
	"github.com/tadasi/portfolio/infrastructure/mysql"
	"github.com/tadasi/portfolio/infrastructure/mysql/tables"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type TodoRepository struct {
}

func NewTodoRepository(i *do.Injector) (models.TodoRepository, error) {
	return &TodoRepository{}, nil
}

func (r *TodoRepository) Find(ctx context.Context, id string) (*models.Todo, error) {
	db, err := mysql.Open()
	if err != nil {
		return nil, err
	}
	todo, err := tables.Todos(
		tables.TodoWhere.ID.EQ(id),
	).One(ctx, db)
	if err != nil {
		return nil, err
	}
	return r.convertToModel(todo), nil
}

func (r *TodoRepository) Create(ctx context.Context, todo *models.Todo) error {
	db, err := mysql.Open()
	if err != nil {
		return err
	}
	record := r.convertToTable(todo)
	if err := record.Insert(ctx, db, boil.Infer()); err != nil {
		return err
	}
	return nil
}

func (r *TodoRepository) convertToModel(record *tables.Todo) *models.Todo {
	return &models.Todo{
		ID:          record.ID,
		Content:     record.Content,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
		CompletedAt: record.CompletedAt.Ptr(),
	}
}

func (r *TodoRepository) convertToTable(todo *models.Todo) *tables.Todo {
	return &tables.Todo{
		ID:          todo.ID,
		Content:     todo.Content,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
		CompletedAt: null.TimeFromPtr(todo.CompletedAt),
	}
}
