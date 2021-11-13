package database

import (
	"context"
	"github.com/tranngoclam/go-msvc/internal/model"
)

type Task struct {
	*DB
}

func (db *Task) Put(ctx context.Context, t *model.Task) error {
	return db.db.WithContext(ctx).Create(t).Error
}

func (db *Task) Get(ctx context.Context, id uint64) (*model.Task, error) {
	t := &model.Task{}
	err := db.db.WithContext(ctx).First(t, id).Error
	return t, err
}
