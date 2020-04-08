CREATE TABLE employee (
	id serial NOT NULL,
	family_name varchar(100) NOT NULL,
	first_name varchar(100) NOT NULL,
	patronymic varchar(100) NULL,
	employee_public_id uuid NOT NULL,
	CONSTRAINT employee_pkey PRIMARY KEY (id)
);