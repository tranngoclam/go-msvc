package database

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/tranngoclam/go-msvc/internal/config"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	cfg, err := config.New(ctx)
	require.NoError(t, err)
	require.NotNil(t, cfg)

	db, err := New(ctx, cfg.Database())
	require.NoError(t, err)
	require.NotNil(t, db)
}
