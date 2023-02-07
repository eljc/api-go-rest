create table products(
    id serial primary key,
    name varchar,
    price decimal
);

INSERT INTO public.products(
	id, name, price)
	VALUES ('Grey blouse', 99.9);

INSERT INTO public.products(
	id, name, price)
	VALUES ('Black jacket', 179);

INSERT INTO public.products(
	id, name, price)
	VALUES ('Sweatshirt', 139.9);