
create table flights (
    id SERIAL PRIMARY KEY,
    name VARCHAR(20) NOT NULL 
);

CREATE TABLE trips (
  id SERIAL PRIMARY KEY,
  flight_id INTEGER NOT NULL,
  name VARCHAR(20) NOT NULL,
  FOREIGN KEY (flight_id) REFERENCES flights (id)
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE seats (
  id SERIAL PRIMARY KEY,
  trip_id INTEGER NOT NULL,
  user_id INTEGER,
  name VARCHAR(10) NOT NULL,
  FOREIGN KEY (trip_id) REFERENCES trips (id),
  FOREIGN KEY (user_id) REFERENCES users (id)
);

insert into flights (name) values ('spicejet');

insert into users (name)
SELECT 'Name ' || num || ' ' || substring(md5(random()::text), 1, 10)
FROM generate_series(1, 120) AS num;


with last_flight AS (
    select id, name from flights order by id desc limit 1
)
insert into trips (flight_id, name) select id, 'trip' || name from last_flight;

with last_trip AS (
    select id, name from trips order by id desc limit 1
)
insert into seats (trip_id, name)  SELECT last_trip.id, num || chr(64 + letter_num)
FROM generate_series(1, 20) AS num
CROSS JOIN generate_series(1, 6) AS letter_num
CROSS JOIN last_trip
ORDER BY num, letter_num;
