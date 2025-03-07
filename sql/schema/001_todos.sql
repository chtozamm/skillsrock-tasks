-- +goose Up
CREATE TABLE tasks(
	id SERIAL PRIMARY KEY,
	title TEXT NOT NULL,
	description TEXT,
	status TEXT CHECK (status IN ('new', 'in_progress', 'done')) DEFAULT 'new',
	created_at TIMESTAMP DEFAULT NOW(),
	updated_at TIMESTAMP DEFAULT NOW()
);

-- +goose Down
DROP TABLE tasks;
