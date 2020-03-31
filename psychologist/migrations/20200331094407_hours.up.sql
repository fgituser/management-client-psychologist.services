-- public.hours definition

-- Drop table

-- DROP TABLE public.hours;

CREATE TABLE public.hours (
	id serial NOT NULL,
	start_time time NOT NULL,
	CONSTRAINT hours_pkey PRIMARY KEY (id)
);