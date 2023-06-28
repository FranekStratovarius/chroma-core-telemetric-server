CREATE TABLE ragequits (
	id			SERIAL PRIMARY KEY NOT NULL,
	level		INT NOT NULL,
	time		INT NOT NULL,
	positionx	INT NOT NULL,
	positiony	INT NOT NULL
);

CREATE TABLE wins (
	id			SERIAL PRIMARY KEY NOT NULL,
	level		INT NOT NULL,
	time		INT NOT NULL
);

CREATE TABLE deaths (
	id			SERIAL PRIMARY KEY NOT NULL,
	level		INT NOT NULL,
	time		INT NOT NULL,
	positionx	INT NOT NULL,
	positiony	INT NOT NULL
);