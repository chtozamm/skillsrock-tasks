-- name: GetTasks :many
SELECT * FROM tasks ORDER BY created_at;

-- name: GetTaskByID :one
SELECT * FROM tasks WHERE id = $1;

-- name: UpdateTask :one
UPDATE tasks SET title = $1, description = $2, status = $3, updated_at = NOW()
WHERE id = $4
RETURNING *;

-- name: DeleteTask :exec
DELETE FROM tasks WHERE id = $1;

-- name: CreateTask :one
INSERT INTO tasks (title, description)
VALUES (
	$1,
	$2
)
RETURNING *;
