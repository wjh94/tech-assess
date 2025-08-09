CREATE TABLE user_activity_logs (
    id bigserial primary key,
    user_id bigint not null references users(id),
    activity text not null,
    created_at timestamp not null default current_timestamp
);

-- This migration file is used to create the user_activity_logs table.