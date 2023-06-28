CREATE TABLE ragequits (
	id			INT AUTO_INCREMENT NOT NULL,
	level		INT NOT NULL,
	time		INT NOT NULL,
	positionx	INT NOT NULL,
	positiony	INT NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE wins (
	id		INT AUTO_INCREMENT NOT NULL,
	level	INT NOT NULL,
	time	INT NOT NULL,
	PRIMARY KEY (`id`)
);

CREATE TABLE deaths (
	id			INT AUTO_INCREMENT NOT NULL,
	level		INT NOT NULL,
	time		INT NOT NULL,
	positionx	INT NOT NULL,
	positiony	INT NOT NULL,
	PRIMARY KEY (`id`)
);