\c todos

DROP TABLE IF EXISTS todos;
CREATE TABLE todos (
    todo_id varchar(255) PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT false
);
