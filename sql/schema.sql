CREATE TABLE Todo (
    id INTEGER primary key,
    content TEXT NOT NULL,
    done BOOLEAN NOT NULL DEFAULT FALSE
);