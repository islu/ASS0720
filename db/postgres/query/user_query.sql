
/*
    task
*/

-- Create task
-- name: CreateTask :one
INSERT INTO task (
    task_group_no, task_name, task_desc, start_time, end_time
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- Get task
-- name: GetTask :one
SELECT * FROM task
WHERE seqno = $1;

-- List task
-- name: ListTask :many
SELECT * FROM task
ORDER BY start_time;

-- List task by task_group_no
-- name: ListTaskByGroupNo :many
SELECT * FROM task
WHERE task_group_no = $1
ORDER BY start_time;

/*
    user_task
*/

-- Create user task
-- name: CreateUserTask :one
INSERT INTO user_task (
    task_seqno, wallet_address, total_amount, point, status, create_time, update_time
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- List user task & task
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