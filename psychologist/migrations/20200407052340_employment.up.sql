CREATE TABLE employment (
	id serial NOT NULL,
	client_id int4 NOT NULL,
	sсhedule_id int4 NOT NULL,
	CONSTRAINT employment_pkey PRIMARY KEY (id)
);


-- public.employment foreign keys

ALTER TABLE employment ADD CONSTRAINT employment_client_id_fkey FOREIGN KEY (client_id) REFERENCES clients(id) ON UPDATE RESTRICT ON DELETE RESTRICT;
ALTER TABLE employment ADD CONSTRAINT employment_shedule_id_fkey FOREIGN KEY ("sсhedule_id") REFERENCES "sсhedule"(id) ON UPDATE RESTRICT ON DELETE RESTRICT;