INSERT INTO Station(name) VALUES
('Томск-1'), ('Тайга-1'), ('Юрга-1'), ('Болотная'), ('Новосибирск');

-- YYYY-MM-DD HH:MM:SS
INSERT INTO Route(sending, arrival, from_station_id, to_station_id, distance) VALUES
(datetime('2026-05-29 12:00'), datetime('2026-05-29 12:00', '3 hours'), 1, 5, 300),
(datetime('2026-05-29 13:00'), datetime('2026-05-29 13:00', '3 hours'), 1, 5, 300),
(datetime('2026-05-30 12:00'), datetime('2026-05-30 12:00', '3 hours'), 1, 5, 300);

INSERT INTO List_stations(route_id, station_id, stopping_time) VALUES
(1, 2, '00:10:00'), (1, 3, '00:05:00'), (1, 4, '00:05:00');
