CREATE TABLE public.clients (
	id serial NOT NULL,
	is_active bool NOT NULL,
	client_public_id uuid NOT NULL,
	family_name varchar(100) NOT NULL,
	first_name varchar(100) NOT NULL,
	patronymic varchar(100) NULL,
	employee_id int4 NOT NULL,
	CONSTRAINT clients_pkey PRIMARY KEY (id)
);


-- clients foreign keys

ALTER TABLE clients ADD CONSTRAINT clients_employee_id_fkey FOREIGN KEY (employee_id) REFERENCES employee(id) ON UPDATE RESTRICT ON DELETE RESTRICT;