CREATE TABLE article_versions (
    id bigserial primary key,
    article_id bigint not null references articles(id) on delete cascade,
    content text not null,
    version_number integer not null,
    status text not null check (status in ('draft', 'published', 'archived')),
    created_at timestamp not null default current_timestamp
);