CREATE TABLE IF NOT EXISTS "product_images"
(
    "product_id" bigint not null,
    "image_url" varchar(255) not null,
    primary key ("product_id", "image_url"),
    constraint "fk_product_id" foreign key ("product_id") references "products" ("id") on delete cascade
);