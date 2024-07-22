package postgres

import (
	"context"

	psqlc "github.com/islu/HW0720/internal/adapter/repository/postgres/postgres_sqlc"
	"github.com/islu/HW0720/internal/domain/user"
)

func (r *PostgresRepository) ListUserTask_Join(ctx context.Context, walletAddress string) ([]user.UserTask, error) {

	q := psqlc.New(r.connPool)

	list, err := q.ListUserTask_Join(ctx, walletAddress)
	if err != nil {
		return nil, err
	}

	var result []user.UserTask
	for _, task := range list {
		result = append(result, user.UserTask{
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
