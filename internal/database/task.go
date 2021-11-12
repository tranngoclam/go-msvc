package database

import (
	"context"
	"github.com/google/uuid"
	"github.com/tranngoclam/go-msvc/internal/model"
)

type TaskDB struct {
	*DB
}

func (db *TaskDB) Create(ctx context.Context, t *model.Task) error {
	return db.db.WithContext(ctx).Create(t).Error
}

func (db *TaskDB) Update(ctx context.Context, id uuid.UUID, t map[string]interface{}) error {
	return db.db.WithContext(ctx).Model(model.Task{ID: id}).Updates(t).Error
}

func (db *TaskDB) Get(ctx context.Context, id uuid.UUID) (*model.Task, error) {
	t := &model.Task{}
	err := db.db.WithContext(ctx).First(t, id).Error
	return t, err
}
