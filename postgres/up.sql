drop table if exists material;
create table material
(
    id           serial primary key,
    title        text,
    description  text,
    ref          text,
    date_created timestamp with time zone default now()
);