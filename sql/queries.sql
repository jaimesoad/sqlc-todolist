-- name: CreateTodo :execresult
INSERT INTO Todo (
    content
) VALUES (?);

-- name: TodoById :one
SELECT * FROM Todo
WHERE id = ? LIMIT 1;

-- name: LastTenTodos :many
SELECT * FROM (SELECT * FROM Todo
ORDER BY id DESC LIMIT 10) as var1
ORDER BY id;

-- name: ToggleTodo :execresult
UPDATE Todo
SET done = NOT done
WHERE id = ?;

-- name: ChangeName :execresult
UPDATE Todo
SET content = ?
WHERE id = ?;

-- name: DeleteTodo :execresult
DELETE FROM Todo
WHERE id = ?;