package postgres

import (
	"context"

	psqlc "github.com/islu/HW0720/internal/adapter/repository/postgres/postgres_sqlc"
	"github.com/islu/HW0720/internal/domain/user"
)

// Create user task
func (r *PostgresRepository) CreateUserTask(ctx context.Context, params user.UserTask) (*user.UserTask, error) {

	tx, err := r.connPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	q := psqlc.New(r.connPool)
	qtx := q.WithTx(tx)

	userTask, err := qtx.CreateUserTask(ctx, psqlc.CreateUserTaskParams{
		TaskSeqno:     int32(params.TaskSeqno),
		WalletAddress: params.WalletAddress,
		Point:         int32(params.Points),
		TotalAmount:   params.TotalAmount,
		Status:        params.Status,
		CreateTime:    params.CreateTime,
		UpdateTime:    params.UpdateTime,
	})
	if err != nil {
		return nil, err
	}

	// Return the created task
	task, err := qtx.GetTask(ctx, int32(userTask.TaskSeqno))
	if err != nil {
		return nil, err
	}

	result := buildUserTask(task, userTask)
	return &result, tx.Commit(ctx)
}

// List user task
func (r *PostgresRepository) ListUserTask_Join(ctx context.Context, walletAddress string) ([]user.UserTask, error) {

	q := psqlc.New(r.connPool)

	list, err := q.ListUserTask_Join(ctx, walletAddress)
	if err != nil {
		return nil, err
	}

	var result []user.UserTask
	for _, task := range list {
		result = append(result, user.UserTask{
			TaskSeqno:       int(task.TaskSeqno),
			TaskName:        task.TaskName.String,
			TaskDescription: task.TaskDesc.String,
			TaskStartTime:   task.StartTime.Time,
			TaskEndTime:     task.EndTime.Time,
			WalletAddress:   task.WalletAddress,
			Points:          int(task.Point),
			TotalAmount:     task.TotalAmount,
			Status:          task.Status,
			CreateTime:      task.CreateTime,
			UpdateTime:      task.UpdateTime,
		})
	}
	return result, nil
}

// Update user task status
func (r *PostgresRepository) UpdateUserTask(ctx context.Context, walletAddress string, params user.UserTask) (*user.UserTask, error) {

	tx, err := r.connPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	q := psqlc.New(r.connPool)
	qtx := q.WithTx(tx)

	userTask, err := qtx.UpdateUserTask(ctx, psqlc.UpdateUserTaskParams{
		TaskSeqno:     int32(params.TaskSeqno),
		WalletAddress: walletAddress,
		Point:         int32(params.Points),
		TotalAmount:   params.TotalAmount,
		Status:        params.Status,
		UpdateTime:    params.UpdateTime,
	})
	if err != nil {
		return nil, err
	}

	// Return the created task
	task, err := qtx.GetTask(ctx, int32(userTask.TaskSeqno))
	if err != nil {
		return nil, err
	}

	result := buildUserTask(task, userTask)
	return &result, tx.Commit(ctx)
}

func buildUserTask(task psqlc.Task, userTask psqlc.UserTask) user.UserTask {
	return user.UserTask{
		TaskName:        task.TaskName,
		TaskDescription: task.TaskDesc,
		TaskStartTime:   task.StartTime,
		TaskEndTime:     task.EndTime,
		WalletAddress:   userTask.WalletAddress,
		Points:          int(userTask.Point),
		TotalAmount:     userTask.TotalAmount,
		Status:          userTask.Status,
		CreateTime:      userTask.CreateTime,
		UpdateTime:      userTask.UpdateTime,
	}
}
