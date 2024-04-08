create table banners
(
    id         serial primary key,
    content       varchar(255) not null,
    is_active boolean,
    created_at timestamp    not null default now(),
    updated_at timestamp    not null default now()
);

create table banners_relations
(
    id         serial primary key,
    tag_id    integer       not null,
    feature_id integer       not null,
    banner_id integer not null,
    created_at timestamp not null default now(),
    updated_at timestamp    not null default now(),
    foreign key (banner_id) references banners (id)
);