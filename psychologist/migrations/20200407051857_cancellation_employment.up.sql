CREATE TABLE cancellation_employment (
	id serial NOT NULL,
	employment_id int4 NOT NULL,
	datetime timestamp NULL DEFAULT timezone('utc'::text, now()),
	CONSTRAINT cancellatioon_employment_pkey PRIMARY KEY (id)
);


-- cancellation_employment foreign keys

ALTER TABLE cancellation_employment ADD CONSTRAINT cancellation_employment_employment_id_fkey FOREIGN KEY (employment_id) REFERENCES employment(id) ON UPDATE CASCADE ON DELETE CASCADE;