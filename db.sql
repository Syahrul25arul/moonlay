CREATE TYPE status AS ENUM ('Y', 'N');
-- create table users with relation customers
CREATE TABLE customers (
    id serial not null,
    customer_id varchar (20) not null,
    customer_name varchar (50) not null,
    status status default 'N',
    CONSTRAINT users_pkey PRIMARY KEY (id)
);


CREATE TABLE products (
    id serial not null,
    product_id varchar(20) not null,
    product_name varchar(50) not null,
    currency_code varchar(10),
    status status default 'N',
    CONSTRAINT token_pkey PRIMARY KEY (id)
);

CREATE TABLE transactions (
    id serial not null,
    transaction_id varchar(20) not null,
    transaction_uuid varchar(100),
    rel_uuid varchar(100),
    buyer_id varchar(20),
    seller_id varchar(20),
    product_id varchar(20),
    price integer,
    volume integer,
    value bigint,
    transaction_date timestamp,
    entry_date timestamp,
    confirm_date timestamp,
    complete_data_buyer timestamp,
    complete_data_seller timestamp,
    buy_sell varchar(5),
    is_amend status,
    is_cancel status,
    confirm_status varchar(20),
    complete_status_buyer varchar(20),
    complete_status_seller varchar(20),
    status status,
    CONSTRAINT token_pkey PRIMARY KEY (id)
)

-- -- create table customers
-- CREATE TABLE customers (
--     customer_id serial primary key,
--     name varchar (25) not null,
--     date_of_birth date not null,
--     zip_code varchar(10) not null,
--     status status default 'inactive',
--     created_at timestamp default now()
-- );

-- -- create index for searching with where clause
-- CREATE INDEX customers_zip_code_and_status ON customers(zip_code, status);



-- -- create table category
-- CREATE TABLE categories (
--     category_id serial not null,
--     category_name varchar(20) not null,
--     CONSTRAINT category_pkey PRIMARY KEY (category_id)
-- );

-- -- create table product
-- CREATE TABLE products (
--     product_id serial not null,
--     product_name varchar(255) not null,
--     category_id integer references categories(category_id),
--     price decimal DEFAULT 0,
--     stock integer DEFAULT 0,
--     product_description text,
--     CONSTRAINT product_pkey PRIMARY KEY (product_id)
-- );

-- -- create index for searching where cluase by category and stock
-- CREATE INDEX product_category_status ON products(category_id,stock);

-- ALTER TABLE products ADD CONSTRAINT price_check check (price >= 0);
-- ALTER TABLE products ADD CONSTRAINT stock_check check (stock >= 0);

-- -- create table images
-- CREATE TABLE images (
--     product_id integer references products(product_id),
--     image_url varchar(255)
-- );

