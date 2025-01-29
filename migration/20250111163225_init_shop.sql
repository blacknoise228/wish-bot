-- +goose Up
-- +goose StatementBegin
create table if not exists "shop" (
    id uuid not null primary key default gen_random_uuid(),
    name varchar(255) not null,
    description varchar(255),
    image varchar(255) not null,
    token varchar(255) not null unique,
    created_at timestamp default now(),
    updated_at timestamp
);

create table if not exists "shop_admins" (
    admin_id bigint not null primary key unique,
    shop_id uuid not null,
    foreign key (shop_id) references "shop" (id)
);

create table if not exists "dim_product_category" (
    id integer not null primary key,
    category_code varchar(255) not null,
    category_name varchar(255) not null
);

insert into "dim_product_category" (id, category_code, category_name) values (1, 'flowers', 'Цветы');
insert into "dim_product_category" (id, category_code, category_name) values (2, 'clothes', 'Одежда');
insert into "dim_product_category" (id, category_code, category_name) values (3, 'electronics', 'Электроника');
insert into "dim_product_category" (id, category_code, category_name) values (4, 'toys', 'Игрушки');
insert into "dim_product_category" (id, category_code, category_name) values (5, 'accessories', 'Аксессуары');

create table if not exists "dim_product_status" (
    id integer not null primary key,
    status_code varchar(255) not null,
    status_name varchar(255) not null
);

insert into "dim_product_status" (id, status_code, status_name) values (1, 'in_stock', 'В наличии');
insert into "dim_product_status" (id, status_code, status_name) values (2, 'out_of_stock', 'Нет в наличии');

create table if not exists "product" (
    id uuid not null primary key unique default gen_random_uuid(),
    name varchar(255) not null,
    price double precision not null,
    description varchar(255) not null,
    image varchar(255) not null,
    category_id integer not null,
    status integer not null,
    shop_id uuid not null,
    admin_id bigint not null,
    created_at timestamp default now(),
    updated_at timestamp,
    foreign key (shop_id) references "shop" (id),
    foreign key (category_id) references "dim_product_category" (id),
    foreign key (status) references "dim_product_status" (id)
);

create table if not exists "dim_order_status" (
    id integer not null primary key,
    status_code varchar(255) not null,
    status_name varchar(255) not null
);

insert into "dim_order_status" (id, status_code, status_name) values (1, 'in_processing', 'В обработке');
insert into "dim_order_status" (id, status_code, status_name) values (2, 'paid', 'Оплачен');
insert into "dim_order_status" (id, status_code, status_name) values (3, 'delivered', 'Доставлен');
insert into "dim_order_status" (id, status_code, status_name) values (4, 'canceled', 'Отменен');
insert into "dim_order_status" (id, status_code, status_name) values (5, 'out_of_stock', 'Нет в наличии');

create table if not exists "orders" (
    id uuid not null primary key unique default gen_random_uuid(),
    price double precision not null,
    status integer not null,
    customer_id bigint not null,
    customer_login varchar(255) not null,
    consignee_id bigint not null,
    product_id uuid not null,
    admin_id bigint not null,
    shop_id uuid not null,
    created_at timestamp default now(),
    updated_at timestamp,
    foreign key (product_id) references "product" (id),
    foreign key (consignee_id) references "users" (chat_id),
    foreign key (status) references "dim_order_status" (id),
    foreign key (admin_id) references "shop_admins" (admin_id)
);

create table if not exists "user_info" (
    chat_id bigint not null primary key unique,
    address varchar(255) not null,
    phone varchar(255) not null,
    name varchar(255) not null,
    description varchar(255) not null,
    created_at timestamp default now(),
    updated_at timestamp,
    foreign key (chat_id) references "users" (chat_id)
);

alter table if exists "wish" add foreign key ("product_id") references "product" ("id");

alter table if exists "wish" add constraint product_id_unique unique (chat_id, product_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists "user_info";    
drop table if exists "orders";
drop table if exists "product";
drop table if exists "shop_admins";
drop table if exists "shop";
drop table if exists "dim_product_status";
drop table if exists "dim_product_category";
drop table if exists "dim_order_status";

-- +goose StatementEnd
