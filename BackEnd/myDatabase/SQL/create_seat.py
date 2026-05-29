insert_query = "INSERT INTO Seat(number, carriage_id) VALUES "
num_seat = int(input("Количество мест: "))
id_seat = int(input("id вагона: "))
values_str = ""
for i in range(1, num_seat):
	values_str += f"({i}, {id_seat}), "
values_str += f"({num_seat}, {id_seat});"

print(insert_query + values_str)
