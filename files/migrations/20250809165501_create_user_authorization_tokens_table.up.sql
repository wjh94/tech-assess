CREATE TABLE user_authorization_tokens (
    id bigserial primary key,
    user_id bigint not null references users(id),
    token text not null unique,
    created_at timestamp not null default current_timestamp,
    expires_at timestamp not null
);