package database

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/tranngoclam/go-msvc/internal/config"
	"github.com/tranngoclam/go-msvc/internal/model"
	"testing"
	"time"
)

func TestTask_PutAndGet(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	cfg, err := config.New(ctx)
	require.NoError(t, err)
	require.NotNil(t, cfg)

	db, err := New(ctx, cfg.Database())
	require.NoError(t, err)
	require.NotNil(t, db)

	taskDB := &Task{DB: db}
	want := &model.Task{
		Description: "my first task",
		Priority:    model.TaskPriorityMedium,
		Status:      model.TaskStatusToDo,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err = taskDB.Put(ctx, want)
	require.NoError(t, err)

	got, err := taskDB.Get(ctx, want.ID)
	require.NoError(t, err)
	require.NotNil(t, got)
	require.Equal(t, want.Description, got.Description)
	require.Equal(t, want.Priority, got.Priority)
	require.Equal(t, want.Status, got.Status)
}
