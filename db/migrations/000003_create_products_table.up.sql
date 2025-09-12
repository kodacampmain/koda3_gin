-- public.products definition
-- Drop table
-- DROP TABLE public.products;
CREATE TABLE public.products (
    id serial4 NOT NULL,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP NULL,
    updated_at timestamp NULL,
    "name" varchar NULL,
    promo_id int4 NULL,
    price int4 DEFAULT 0 NOT NULL,
    CONSTRAINT products_pkey PRIMARY KEY (id),
    CONSTRAINT unique_product_name UNIQUE (name)
);
-- public.products foreign keys
ALTER TABLE public.products
ADD CONSTRAINT products_promo_id_fkey FOREIGN KEY (promo_id) REFERENCES public.promos(id);