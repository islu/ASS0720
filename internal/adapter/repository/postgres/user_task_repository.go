package postgres

import (
	"context"

	psqlc "github.com/islu/ASS0720/internal/adapter/repository/postgres/postgres_sqlc"
	"github.com/islu/ASS0720/internal/domain/user"
)

func (r *PostgresRepository) GetUserTaskList(ctx context.Context) ([]user.UserTask, error) {

	q := psqlc.New(r.connPool)

	list, err := q.ListUserTask(ctx)
	if err != nil {
		return nil, err
	}

	var result []user.UserTask
	for _, task := range list {
		result = append(result, user.UserTask{
			WalletAddress: task.WalletAddress,
			Points:        int(task.Point),
			Amount:        task.Amount,
		})
	}
	return result, nil
}
