package models

import (
	"context"
	"time"
)

type Todo struct {
	ID          string     // TODO ID
	Content     string     // TODO 内容
	CreatedAt   time.Time  // TODO 作成日時
	UpdatedAt   time.Time  // TODO 更新日時
	CompletedAt *time.Time // TODO 完了日時
}

type TodoRepository interface {
	Find(ctx context.Context, id string) (*Todo, error)
}
