SELECT
	r.id AS route_id,
	r.sending,
	r.arrival,
	r.distance,
	s_from.name AS from_station,
	s_to.name AS to_station,
	sor.id AS station_on_route_id,
	s_middle.name AS intermediate_station,
	sor.arrival AS intermediate_arrival
FROM Route r
LEFT JOIN Station s_from ON r.from_station_id = s_from.id
LEFT JOIN Station s_to ON r.to_station_id = s_to.id
LEFT JOIN Stations_on_route sor ON r.id = sor.route_id
LEFT JOIN Station s_middle ON sor.station_id = s_middle.id
WHERE
	s_from.name = ?
	AND s_to.name = ?
	AND DATE(r.sending) >= ?
ORDER BY r.id, sor.arrival;