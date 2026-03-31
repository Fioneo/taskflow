CREATE SCHEMA taskflow;

CREATE TABLE taskflow.users(
    id SERIAL PRIMARY KEY,
    version int NOT NULL DEFAULT 1,
    full_name varchar(100) NOT NULL CHECK(char_length(full_name) BETWEEN 3 AND 100),
    phone_number varchar(15) CHECK(
        phone_number ~ '^\+[0-9]+$'
        AND
        char_length(phone_number) BETWEEN 10 AND 15
    )
);
CREATE TABLE taskflow.tasks(
    id SERIAL PRIMARY KEY,
    version int NOT NULL DEFAULT 1,
    title varchar(100) NOT NULL CHECK(char_length(title) BETWEEN 1 AND 30),
    description varchar(100) CHECK(char_length(description) BETWEEN 1 AND 100),
    completed BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ ,

    CHECK(
        (completed=FALSE AND completed_at IS NULL)
        OR
        (completed=TRUE AND completed_at IS NOT NULL AND completed_at >= created_at)
    ),

    author_id INTEGER NOT NULL REFERENCES taskflow.users(id)
);