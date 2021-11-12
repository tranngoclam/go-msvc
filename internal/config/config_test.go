package config

import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig_New(t *testing.T) {
	user, password := uuid.New().String(), uuid.New().String()

	t.Setenv("DB_HOST", "mysql.cloud.example.com")
	t.Setenv("DB_PORT", "3306")
	t.Setenv("DB_USER", user)
	t.Setenv("DB_PASSWORD", password)
	t.Setenv("DB_NAME", "go-msvc")
	t.Setenv("PPROF_PORT", "9207")

	cfg, err := New(context.Background())
	require.NoError(t, err)
	require.NotNil(t, cfg)

	db := cfg.Database()
	require.NotNil(t, db)
	require.Equal(t, "mysql.cloud.example.com", db.Host)
	require.Equal(t, 3306, db.Port)
	require.Equal(t, user, db.User)
	require.Equal(t, password, db.Password)
	require.Equal(t, "go-msvc", db.Name)

	tm := cfg.Telemetry()
	require.NotNil(t, tm)
	require.Equal(t, 9207, tm.PprofPort)
}
