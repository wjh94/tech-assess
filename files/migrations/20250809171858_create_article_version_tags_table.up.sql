CREATE TABLE article_version_tags (
    id bigserial primary key,
    article_version_id bigint not null references article_versions(id) on delete cascade,
    tag_id bigint not null references tags(id) on delete cascade,
    created_at timestamp not null default current_timestamp
);