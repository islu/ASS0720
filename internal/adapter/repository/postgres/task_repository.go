package postgres

import (
	"context"

	psqlc "github.com/islu/ASS0720/internal/adapter/repository/postgres/postgres_sqlc"
	"github.com/islu/ASS0720/internal/domain/user"
)

// Create a new task
func (r *PostgresRepository) CreateTask(ctx context.Context, params user.Task) (*user.Task, error) {

	tx, err := r.connPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	q := psqlc.New(r.connPool)
	qtx := q.WithTx(tx)

	task, err := qtx.CreateTask(ctx, psqlc.CreateTaskParams{
		TaskGroupNo: int32(params.GroupNo),
		TaskName:    params.Name,
		TaskDesc:    params.Description,
		StartTime:   params.StartTime,
		EndTime:     params.EndTime,
	})

	if err != nil {
		return nil, err
	}

	result := &user.Task{
		GroupNo:     int(task.TaskGroupNo),
		Name:        task.TaskName,
		Description: task.TaskDesc,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,
	}
	return result, tx.Commit(ctx)
}

// Get task list by group no
func (r *PostgresRepository) ListTaskByGroupNo(ctx context.Context, groupNo int) ([]user.Task, error) {

	q := psqlc.New(r.connPool)

	list, err := q.ListTaskByGroupNo(ctx, int32(groupNo))
	if err != nil {
		return nil, err
	}

	var result []user.Task
	for _, task := range list {
		result = append(result, user.Task{
			GroupNo:     int(task.TaskGroupNo),
			Name:        task.TaskName,
			Description: task.TaskDesc,
			StartTime:   task.StartTime,
			EndTime:     task.EndTime,
		})
	}

	return result, nil
}
