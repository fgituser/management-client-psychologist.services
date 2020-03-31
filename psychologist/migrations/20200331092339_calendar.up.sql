-- public.calendar definition

-- Drop table

-- DROP TABLE public.calendar;

CREATE TABLE public.calendar (
	day_id date NOT NULL,
	"year" int2 NOT NULL,
	"month" int2 NOT NULL,
	"day" int2 NOT NULL,
	quarter int2 NOT NULL,
	CONSTRAINT calendar_pkey PRIMARY KEY (day_id),
	CONSTRAINT con_month CHECK (((month >= 1) AND (month <= 31)))
)