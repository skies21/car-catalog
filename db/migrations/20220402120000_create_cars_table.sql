-- +goose Up
CREATE SCHEMA IF NOT EXISTS car_schema;
SET search_path TO car_schema;

CREATE TABLE IF NOT EXISTS goose_db_version (
    id serial primary key,
    version_id bigint not null,
    is_applied bool not null default false,
    tstamp timestamp not null default CURRENT_TIMESTAMP
);


CREATE TABLE IF NOT EXISTS cars (
    id SERIAL PRIMARY KEY,
    reg_num VARCHAR(20) NOT NULL,
    mark VARCHAR(100),
    model VARCHAR(100),
    year INTEGER,
    owner_name VARCHAR(100),
    owner_surname VARCHAR(100),
    owner_patronymic VARCHAR(100)
);

-- +goose Down
DROP TABLE IF EXISTS cars;
