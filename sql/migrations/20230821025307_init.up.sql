CREATE TYPE user_status_type_enum AS ENUM ('active', 'locked','deleted');

CREATE TABLE IF NOT EXISTS "users" 
(
    "id" bigserial not null,
    "name" varchar(50),
    "username" varchar unique not null,
    "encrypted_password" varchar not null,
    "email" varchar(50) not null unique,
    "phone_number" varchar(15) unique,
    "avatar" varchar,
    "role" smallint default 1,  
    "status" user_status_type_enum default 'active' not null,
    "created_at" timestamp default (now() at time zone 'UTC'),
    "updated_at" timestamp default (now() at time zone 'UTC'),
    primary key ("id")
);

CREATE TABLE IF NOT EXISTS "addresses" 
(
    "id" bigserial not null,
    "user_id" bigint not null,
    "country" varchar(20),
    "city" varchar(20),
    "home_address" varchar,
    "created_at" timestamp default (now() at time zone 'UTC'),
    "updated_at" timestamp default (now() at time zone 'UTC'),
    primary key ("id"),
    constraint "fk_user_id" foreign key ("user_id") references "users" ("id") on delete cascade
);

CREATE TABLE IF NOT EXISTS "carts"
(
    "id" bigserial not null,
    "user_id" bigint not null,
    primary key ("id"),
    "created_at" timestamp default (now() at time zone 'UTC'),
    "updated_at" timestamp default (now() at time zone 'UTC'),
    constraint "fk_user_id" foreign key ("user_id") references "users" ("id") on delete cascade
);


CREATE TABLE IF NOT EXISTS "categories" 
(
    "name" varchar(50) unique not null,
    "slug" varchar(100) unique not null,
    "created_at" timestamp default (now() at time zone 'UTC'),
    "updated_at" timestamp default (now() at time zone 'UTC'),
    primary key ("name")
);

CREATE TABLE IF NOT EXISTS "brands"
(
    "name" varchar(50) unique not null,
    "slug" varchar(100) unique not null,
    "description" varchar,
    "created_at" timestamp default (now() at time zone 'UTC'),
    "updated_at" timestamp default (now() at time zone 'UTC'),
    primary key ("name")
);

CREATE TABLE IF NOT EXISTS "products"
(
    "id" bigserial not null,
    "name" varchar not null,
    "thumbnail" varchar,
    "brand" varchar(50) not null,
    "created_at" timestamp default (now() at time zone 'UTC'),
    "updated_at" timestamp default (now() at time zone 'UTC'),
    primary key ("id"),
    constraint "fk_brand" foreign key ("brand") references "brands" ("name") on delete cascade
);


CREATE TABLE IF NOT EXISTS "color_products"
(
    "color" varchar not null,
    "product_id" bigint not null,
    "image" varchar,
    "price" int not null default 0,
    "quantity" smallint not null default 0,
    "sold" smallint not null default 0,
    "created_at" timestamp default (now() at time zone 'UTC'),
    "updated_at" timestamp default (now() at time zone 'UTC'),
    primary key ("color","product_id"),
    constraint "fk_product_id" foreign key ("product_id") references "products" ("id") on delete cascade
);

CREATE TABLE IF NOT EXISTS "products_categories"
(
    "product_id" bigint not null,
    "category_name" varchar(50) not null,
    primary key ("product_id", "category_name"),
    constraint "fk_product_id" foreign key ("product_id") references "products" ("id") on delete cascade,
    constraint "fk_category_name" foreign key ("category_name") references "categories" ("name") on delete cascade
);

CREATE TABLE IF NOT EXISTS "properties"
(
    "id" bigserial not null,
    "product_id" bigint not null,
    "name" varchar not null,
    "value" varchar not null,
    primary key ("id"),
    constraint "fk_product_id" foreign key ("product_id") references "products" ("id") on delete cascade
);

CREATE TABLE IF NOT EXISTS "carts_products"
(
    "cart_id" bigint not null,
    "product_id" bigint not null,
    "quantity" int not null default 0,
    "created_at" timestamp default (now() at time zone 'UTC'),
    "updated_at" timestamp default (now() at time zone 'UTC'),
    primary key ("cart_id", "product_id"),
    constraint "fk_cart_id" foreign key ("cart_id") references "carts" ("id") on delete cascade,
    constraint "fk_product_id" foreign key ("product_id") references "products" ("id") on delete cascade
);
