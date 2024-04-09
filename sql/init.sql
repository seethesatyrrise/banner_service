create table banners
(
    banner_id         serial primary key,
    content     jsonb not null,
    tag_ids    bigint[]       not null,
    feature_id bigint       not null,
    is_active boolean,
    created_at timestamp    not null default now(),
    updated_at timestamp    not null default now()
);