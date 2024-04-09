create table banners
(
    id         serial primary key,
    content     jsonb not null,
    tag_ids    integer[]       not null,
    feature_id integer       not null,
    is_active boolean,
    created_at timestamp    not null default now(),
    updated_at timestamp    not null default now()
);