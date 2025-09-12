-- public.students definition
-- Drop table
-- DROP TABLE public.students;
CREATE TABLE public.students (
    id serial4 NOT NULL,
    "name" varchar NULL,
    "password" varchar NULL,
    "role" varchar NULL,
    image varchar(254) NULL,
    CONSTRAINT students_name_key UNIQUE (name),
    CONSTRAINT students_pkey PRIMARY KEY (id)
);