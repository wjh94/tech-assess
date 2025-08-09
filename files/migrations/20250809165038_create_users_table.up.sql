CREATE TABLE users (
    id bigserial primary key,
    "username" text not null unique,
    "email" text not null unique,
    password_hash text not null,
    "role_id" bigint not null references user_roles(id),
    "status" int default 1,
    created_at timestamp not null default current_timestamp
);

-- This migration file is used to create the users table.