-- data from calendar
INSERT INTO calendar (day_id, year, month, day, quarter)
(SELECT ts, 
  EXTRACT(YEAR FROM ts),
  EXTRACT(MONTH FROM ts),
  EXTRACT(DAY FROM ts),
  EXTRACT(QUARTER FROM ts)
  FROM generate_series('2020-03-01'::timestamp, '2021-01-01', '1day'::interval) AS t(ts));

-- data from employee
INSERT INTO   employee (id, family_name, first_name, patronymic, employee_public_id) VALUES (1, 'Иванов', 'Иван', 'Иванович', '75d2cdd6-cf69-44e7-9b28-c47792505d81');
INSERT INTO   employee (id, family_name, first_name, patronymic, employee_public_id) VALUES (2, 'Петров', 'Петр', 'Петрович', '11e195fc-7010-4e50-8a4d-1d43e9c8e5db');
INSERT INTO   employee (id, family_name, first_name, patronymic, employee_public_id) VALUES (3, 'Сидр', 'Разман', 'Усмотнович', 'ccf8444d-ec0d-4b7e-bf0a-06701f48ae1d');

-- data from hours
insert into hours (start_time)
(
SELECT t.h::time
FROM   generate_series(timestamp '2008-03-07 00:00'
                     , timestamp '2008-03-07 23:00'
                     , interval  '1 hour') AS t(h)
);

--data from schedule
INSERT INTO sсhedule (employee_id, calendar_id, hour_id)
(select e.id,
c.day_id , h.id 
	from employee e, calendar c, hours h
	order by e.id 
);

-- data from clients DELETE FIO!
INSERT INTO   clients (id, is_active, client_public_id, employee_id) VALUES (1, true, '48faa486-8e73-4c31-b10f-c7f24c115cda' 1);
INSERT INTO   clients (id, is_active, client_public_id, employee_id) VALUES (2, true, '50faa486-8e73-4c31-b10f-c7f24c115cda', 2);
INSERT INTO   clients (id, is_active, client_public_id, employee_id) VALUES (3, true, '60faa486-8e73-4c31-b10f-c7f24c115cda', 1);

-- data from employment
INSERT INTO   employment (id, client_id, sсhedule_id) VALUES (1, 1, 11);
INSERT INTO   employment (id, client_id, sсhedule_id) VALUES (2, 2, 22);
INSERT INTO   employment (id, client_id, sсhedule_id) VALUES (3, 3, 13);

--data from cancelation_employment
INSERT INTO   cancellation_employment (id, employment_id, datetime) VALUES (1, 1, '2020-04-07 02:48:30.941402');