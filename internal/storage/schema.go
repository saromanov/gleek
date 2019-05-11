package storage

var taskSchema = `
CREATE TABLE tasks (
	id int8 PRIMARY KEY,
	created_at timestampz NOT NULL,
	name text,
	priority int4 NOT NULL,
	start timestampz,
	duration string
);
`
