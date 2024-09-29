create table if not exists tickets
(
    id          serial                   not null primary key,
    name        text                     not null,
    description text                     not null,
    allocation  integer                  not null,
    created_at  timestamp with time zone not null default current_timestamp,
    updated_at  timestamp with time zone
);
