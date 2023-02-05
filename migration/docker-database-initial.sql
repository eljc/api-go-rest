create table products(
    id serial primary key,
    name varchar,
    price decimal
);

INSERT INTO public.products(
	id, name, price)
	VALUES (1, 'Grey blouse', 99.9);

INSERT INTO public.products(
	id, name, price)
	VALUES (2, 'Black jacket', 179);

INSERT INTO public.products(
	id, name, price)
	VALUES (3, 'Sweatshirt', 139.9);