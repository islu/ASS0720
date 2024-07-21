
-- Create task
-- name: CreateTask :one
INSERT INTO task (
    task_group_no, task_name, task_desc, start_time, end_time
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- Get task by task_group_no
-- name: ListTaskByGroupNo :many
SELECT * FROM task
WHERE task_group_no = $1
ORDER BY start_time;

-- Get user task
-- name: ListUserTask_Join :many
SELECT
    t.task_name,
    t.task_desc,
    t.start_time,
    t.end_time,
    ut.wallet_address,
    ut.total_amount,
    ut.point,
    ut.status,
    ut.create_time,
    ut.update_time
FROM user_task ut
LEFT JOIN task t ON ut.task_seqno = t.seqno
WHERE ut.wallet_address = $1
ORDER BY t.start_time desc;