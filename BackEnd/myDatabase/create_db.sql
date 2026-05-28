CREATE TABLE Users (
	id INTEGER PRIMARY KEY,
	login VARCHAR(50) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	email VARCHAR(100) UNIQUE NOT NULL,
	name VARCHAR(50) NOT NULL,
	surname VARCHAR(60) NOT NULL,
	patronymic VARCHAR(60)
);

CREATE TABLE Station(
	id INTEGER PRIMARY KEY,
	name VARCHAR(50) NOT NULL
);

CREATE TABLE Route(
	id INTEGER PRIMARY KEY,
	sending TEXT DEFAULT (datetime('now')),
	arrival TEXT DEFAULT (datetime('now', '6 hour')),
	from_station_id INTEGER NOT NULL REFERENCES Station(id),
	to_station_id INTEGER NOT NULL REFERENCES Station(id),
	distance SMALLINT NOT NULL
);

CREATE TABLE List_stations(
	id INTEGER PRIMARY KEY,
	route_id INTEGER NOT NULL REFERENCES Route(id),
	station_id INTEGER NOT NULL REFERENCES Station(id),
	stopping_time TEXT,
	completed BOOLEAN DEFAULT FALSE
);