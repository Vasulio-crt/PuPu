SELECT r.sending, r.arrival, s_from.name AS from_station, s_to.name AS to_station, r.distance
FROM Route r
LEFT JOIN Station s_from ON r.from_station_id = s_from.id
LEFT JOIN Station s_to ON r.to_station_id = s_to.id
WHERE s_from.name = 'Томск' AND s_to.name = 'Новосибирск' AND DATE(r.sending) >= '2025-06-01'
ORDER BY r.sending