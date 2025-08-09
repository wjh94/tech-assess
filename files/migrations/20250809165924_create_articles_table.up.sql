CREATE TABLE articles (
    id bigserial primary key,
    title text not null,
    author_id bigint not null references users(id),
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp
);
-- This migration file is used to create the articles table.