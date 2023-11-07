-- +goose Up
create table "order_status_outbox" (
    id serial primary key,
    order_id int not null,
    order_status text not null,
    is_delivery boolean default false,
    created_at timestamp default now() not null
);

-- +goose Down
drop table "order_status_outbox";
