-- public.shedule definition

-- Drop table

-- DROP TABLE public.shedule;

CREATE TABLE public.shedule (
	id serial NOT NULL,
	employee_id int4 NOT NULL,
	work_hour bool NOT NULL DEFAULT true,
	employment_id int4 NULL,
	calendar_id date NOT NULL,
	hour_id int4 NOT NULL,
	CONSTRAINT shedule_pkey PRIMARY KEY (id)
);


-- public.shedule foreign keys

ALTER TABLE public.shedule ADD CONSTRAINT shedule_employee_id_fkey FOREIGN KEY (employee_id) REFERENCES employee(id) ON UPDATE RESTRICT ON DELETE RESTRICT;