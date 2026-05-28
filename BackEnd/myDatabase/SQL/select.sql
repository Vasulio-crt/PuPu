SELECT sending, arrival, S1.name AS from_station_name, S2.name AS to_station_name, distance
FROM Route
JOIN Station AS S1 ON Route.from_station_id = S1.id
JOIN Station AS S2 ON Route.to_station_id = S2.id
WHERE S1.name = 'Томск-1'
	AND S2.name = 'Новосибирск'
	AND sending >= datetime('2025-05-29')
LIMIT 15