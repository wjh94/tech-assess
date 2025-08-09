CREATE TABLE user_roles (
    id bigserial primary key,
    "name" text not null default 'User',
    "status" int default 1,
    created_at timestamp not null default current_timestamp
);

-- This migration file is used to create the user_roles table.
-- It should be run before creating the users table.
-- Ensure that the database is set up correctly before running this migration.