package postgres

import (
	"context"

	psqlc "github.com/islu/HW0720/internal/adapter/repository/postgres/postgres_sqlc"
	"github.com/islu/HW0720/internal/domain/user"
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

	result := buildTask(task)

	return &result, tx.Commit(ctx)
}

// List tasks
func (r *PostgresRepository) ListTask(ctx context.Context) ([]user.Task, error) {

	q := psqlc.New(r.connPool)

	list, err := q.ListTask(ctx)
	if err != nil {
		return nil, err
	}

	var result []user.Task
	for _, task := range list {
		t := buildTask(task)
		result = append(result, t)
	}

	return result, nil
}

// List tasks by group number
func (r *PostgresRepository) ListTaskByGroupNo(ctx context.Context, groupNo int) ([]user.Task, error) {

	q := psqlc.New(r.connPool)

	list, err := q.ListTaskByGroupNo(ctx, int32(groupNo))
	if err != nil {
		return nil, err
	}

	var result []user.Task
	for _, task := range list {
		t := buildTask(task)
		result = append(result, t)
	}

	return result, nil
}

func buildTask(task psqlc.Task) user.Task {
	return user.Task{
		Seqno:       int(task.Seqno),
		GroupNo:     int(task.TaskGroupNo),
		Name:        task.TaskName,
		Description: task.TaskDesc,
		StartTime:   task.StartTime,
		EndTime:     task.EndTime,
	}
}
