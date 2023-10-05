-- +goose Up
create table cart (
    id serial primary key,
    user_id int not null,
    sku_id int not null,
    count int not null
);

-- +goose Down
drop table cart;
