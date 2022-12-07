CREATE TYPE status AS ENUM ('Y', 'N');
-- create table users with relation customers
CREATE TABLE customers (
    id serial not null,
    customer_id varchar (20) not null,
    customer_name varchar (50) not null,
    status status default 'N',
    CONSTRAINT customers_pkey PRIMARY KEY (id)
);


CREATE TABLE products (
    id serial not null,
    product_id varchar(20) not null,
    product_name varchar(50) not null,
    currency_code varchar(10),
    status status default 'N',
    CONSTRAINT products_pkey PRIMARY KEY (id)
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
    CONSTRAINT transaction_pkey PRIMARY KEY (id)
);

-- create view datamart1

create view datamart as select t.transaction_id
, t.buyer_id 
, c.customer_name as buyer_name
, t.seller_id 
, c2.customer_name as seller_name
, t.product_id 
, p.product_name 
, p.currency_code as currency
, t.price 
, t.volume 
, t.value 
, t.transaction_date 
, extract(month from t.transaction_date) as transaction_month
, extract(year from t.transaction_date) as transaction_year
, t.entry_date
, extract(month from t.entry_date) as entry_month
, extract(year from t.entry_date) as entry_year
, t.buy_sell 
, t.confirm_status
, t.complete_status_buyer
, t.complete_status_seller 
from transactions t 
left join customers c on c.customer_id = t.buyer_id 
left join customers c2 on c2.customer_id = t.seller_id 
left join products p on p.product_id = t.product_id ;


select * from datamart d;

