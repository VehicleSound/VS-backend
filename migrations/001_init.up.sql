create table users
(
    id             varchar               not null
        constraint users_pk
            primary key,
    login          varchar               not null,
    email          varchar               not null,
    confirmed      boolean default false not null,
    "passwordHash" varchar               not null,
    active         boolean default true
);

create table tags
(
    id    varchar not null
        constraint tags_pk
            primary key,
    title varchar
        constraint tags_uk
            unique
);

create table vehicle_types
(
    id    varchar not null
        constraint vehicle_types_pk
            primary key,
    title varchar
);

create table vehicles
(
    id      varchar not null
        constraint vehicles_pk
            primary key,
    name    varchar,
    type_id varchar
        constraint vehicles_vehicle_types_id_fk
            references public.vehicle_types
);

create table files
(
    id      varchar not null
        constraint files_pk
            primary key,
    ext     varchar,
    created timestamp default now()
);

create table sounds
(
    id              varchar not null
        constraint sounds_pk
            primary key,
    name            varchar not null,
    description     varchar,
    author_id       varchar
        constraint sounds_users_id_fk
            references public.users,
    vehicle_id      varchar
        constraint sounds_vehicles_id_fk
            references public.vehicles,
    sound_file_id   varchar
        constraint sounds_files_id_fk2
            references public.files,
    picture_file_id varchar
        constraint sounds_files_id_fk
            references public.files
);

create table sound_tags
(
    sound_id varchar
        constraint sound_tags_sounds_id_fk
            references public.sounds,
    tag_id   varchar
        constraint sound_tags_tags_id_fk
            references public.tags
);

create table favourites
(
    user_id  varchar not null
        constraint favourites_users_id_fk
            references public.users,
    sound_id varchar not null
        constraint favourites_sounds_id_fk
            references public.sounds,
    constraint favourites_pk
        primary key (user_id, sound_id)
);

insert into vehicle_types (id, title) values('1', 'any');
insert into vehicles (id, name, type_id) values ('default', 'default', '1');