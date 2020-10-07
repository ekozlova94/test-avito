create table ads
(
    id     bigserial          not null
        constraint ads_pkey primary key,
    number varchar(80) unique not null,
    price  varchar(80)        not null
);

create table subscribers
(
    id     bigserial   not null
        constraint subscribers_pkey primary key,
    ads_id integer REFERENCES ads,
    email  varchar(80) not null
);

