package usecase

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/islu/ASS0720/internal/adapter/repository/postgres"
	"github.com/islu/ASS0720/internal/usecase/service/user"
)

type Application struct {
	Params      ApplicationParams
	UserService *user.UserService
}

type ApplicationParams struct {
	// Env
	Environment string

	// Database
	DBUser       string
	DBPassword   string
	DBName       string
	DBSchemaName string
}

func NewApplication(ctx context.Context, param *ApplicationParams) (*Application, error) {

	// Initialize database
	pgRepo, err := initDatabase(ctx, *param)
	if err != nil {
		return nil, err
	}

	// New application
	app := &Application{
		Params: *param,
		UserService: user.NewUserService(ctx, user.UserServiceParam{
			UserTaskRepo: pgRepo,
		}),
	}
	return app, nil
}

/*
	Database
*/

func initDatabase(ctx context.Context, cfg ApplicationParams) (*postgres.PostgresRepository, error) {

	conn, err := connect(cfg)
	if err != nil {
		return nil, err
	}
	pgRepo := postgres.NewPostgresRepository(ctx, conn)
	return pgRepo, nil
}

// Connect postgres
func connect(cfg ApplicationParams) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf("user=%s password=%s database=%s search_path=%s", cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSchemaName)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	dbPool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}
	return dbPool, nil
}
