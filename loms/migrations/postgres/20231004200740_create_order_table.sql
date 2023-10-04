-- +goose Up
create type order_status as enum (
    'new',
    'await_payment',
    'failed',
    'payed',
    'canceled'
);

create table "order" (
    id serial primary key,
    user_id int not null,
    status order_status not null,
    items json not null,
    created_at timestamp default now() not null
);

-- +goose Down
drop table "order";
