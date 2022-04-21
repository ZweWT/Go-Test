CREATE TABLE IF NOT EXISTS todos (
    id serial PRIMARY KEY,
    name text NOT NULL,
    user_id integer references users(id),
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
);