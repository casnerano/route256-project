-- +goose Up
create table "stock" (
    id serial primary key,
    sku_id int not null,
    available int not null,
    reserved int not null default 0
);

-- +goose Down
drop table "stock";
