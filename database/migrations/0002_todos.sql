\c todos;

DROP TABLE IF EXISTS todos;
CREATE TABLE todos (
    todo_id serial PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    completed INTEGER NOT NULL DEFAULT 0
);
