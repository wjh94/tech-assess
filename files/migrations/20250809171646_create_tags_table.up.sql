CREATE TABLE tags (
    id bigserial primary key,
    "name" text not null unique,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);