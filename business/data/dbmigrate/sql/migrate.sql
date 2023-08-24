-- Version: 1.01
-- Description: Create table tasks
CREATE TABLE tasks (
  id            UUID        NOT NULL,
	name          TEXT        NOT NULL,
	description   TEXT        NOT NULL,
	type          TEXT        NOT NULL,
	due_date      TIMESTAMP   NOT NULL,
	date_created  TIMESTAMP   NOT NULL,
	date_updated  TIMESTAMP   NOT NULL,

	PRIMARY KEY (id)
);

