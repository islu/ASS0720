
-- Get user points history for distributed tasks
-- name: ListUserTask :many
SELECT * FROM user_task
ORDER BY wallet_address desc;