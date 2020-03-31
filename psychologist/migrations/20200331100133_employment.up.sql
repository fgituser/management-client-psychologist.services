-- public.employment definition

-- Drop table

-- DROP TABLE public.employment;

CREATE TABLE public.employment (
	id serial NOT NULL,
	client_id int4 NOT NULL,
	employee_id int4 NOT NULL,
	CONSTRAINT employment_pkey PRIMARY KEY (id)
);


-- public.employment foreign keys

ALTER TABLE public.employment ADD CONSTRAINT employment_client_id_fkey FOREIGN KEY (client_id) REFERENCES clients(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE public.employment ADD CONSTRAINT employment_employee_id_fkey FOREIGN KEY (employee_id) REFERENCES shedule(id) ON UPDATE RESTRICT ON DELETE RESTRICT;