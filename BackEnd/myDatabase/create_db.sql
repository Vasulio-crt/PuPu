CREATE TABLE IF NOT EXISTS Users (
	id INTEGER PRIMARY KEY,
	login VARCHAR(50) UNIQUE NOT NULL,
	password VARCHAR(255) NOT NULL,
	email VARCHAR(100) UNIQUE NOT NULL,
	name VARCHAR(50) NOT NULL,
	surname VARCHAR(60) NOT NULL,
	patronymic VARCHAR(60)
);

CREATE TABLE IF NOT EXISTS Station(
	id INTEGER PRIMARY KEY,
	name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS Route(
	id INTEGER PRIMARY KEY,
	sending TEXT DEFAULT (datetime('now')),
	arrival TEXT DEFAULT (datetime('now', '6 hour')),
	from_station_id INTEGER NOT NULL REFERENCES Station(id),
	to_station_id INTEGER NOT NULL REFERENCES Station(id),
	distance SMALLINT NOT NULL
);

CREATE TABLE IF NOT EXISTS Carriage(
	id INTEGER PRIMARY KEY,
	type VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS Seat(
	id INTEGER PRIMARY KEY,
	number SMALLINT NOT NULL,
	carriage_id INTEGER NOT NULL REFERENCES Carriage(id),
	occupied INTEGER NOT NULL DEFAULT 0 REFERENCES Users(id)
);

INSERT INTO Station(name) VALUES
('Томск'), ('Тайга'), ('Болотная'), ('Новосибирск');

-- YYYY-MM-DD HH:MM:SS
INSERT INTO Route(sending, arrival, from_station_id, to_station_id, distance) VALUES
(datetime('2026-06-02 12:00'), datetime('2026-06-02 12:00', '3 hours'), 1, 4, 300),
(datetime('2026-06-02 13:00'), datetime('2026-06-02 13:00', '3 hours'), 1, 4, 300),
(datetime('2026-06-03 12:00'), datetime('2026-06-03 12:00', '3 hours'), 1, 4, 300),
(datetime('2026-06-05 12:00'), datetime('2026-06-05 12:00', '3 hours'), 1, 4, 310);

INSERT INTO Carriage(type) VALUES ('Плацкартный'), ('СВ');

INSERT INTO Seat(number, carriage_id) VALUES
(1, 1), (2, 1), (3, 1), (4, 1), (5, 1), (6, 1),
(7, 1), (8, 1), (9, 1), (10, 1), (11, 1), (12, 1),
(13, 1), (14, 1), (15, 1), (16, 1), (17, 1), (18, 1),
(19, 1), (20, 1), (21, 1), (22, 1), (23, 1), (24, 1),
(25, 1), (26, 1), (27, 1), (28, 1), (29, 1), (30, 1),
(31, 1), (32, 1), (33, 1), (34, 1), (35, 1), (36, 1),
(37, 1), (38, 1), (39, 1), (40, 1), (41, 1), (42, 1),
(43, 1), (44, 1), (45, 1), (46, 1), (47, 1), (48, 1),
(49, 1), (50, 1), (51, 1), (52, 1), (53, 1), (54, 1);

INSERT INTO Seat(number, carriage_id) VALUES
(1, 2), (2, 2), (3, 2), (4, 2), (5, 2), (6, 2),
(7, 2), (8, 2), (9, 2), (10, 2), (11, 2), (12, 2),
(13, 2), (14, 2), (15, 2), (16, 2), (17, 2), (18, 2);