CREATE TABLE sсhedule (
	id serial NOT NULL,
	employee_id int4 NOT NULL,
	work_hour bool NOT NULL DEFAULT true,
	calendar_id date NOT NULL,
	hour_id int4 NOT NULL,
	CONSTRAINT shedule_pkey PRIMARY KEY (id)
)

ALTER TABLE public.sсhedule ADD CONSTRAINT shedule_calendar_id_fkey FOREIGN KEY (calendar_id) REFERENCES calendar(day_id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE public.sсhedule ADD CONSTRAINT shedule_employee_id_fkey FOREIGN KEY (employee_id) REFERENCES employee(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE public.sсhedule ADD CONSTRAINT shedule_hour_id_fkey FOREIGN KEY (hour_id) REFERENCES hours(id) ON UPDATE RESTRICT ON DELETE RESTRICT;