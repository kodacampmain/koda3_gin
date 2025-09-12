-- public.promos definition
-- Drop table
-- DROP TABLE public.promos;
CREATE TABLE public.promos (
    id serial4 NOT NULL,
    "name" varchar NOT NULL,
    discount int4 NOT NULL,
    CONSTRAINT promos_pkey PRIMARY KEY (id)
);