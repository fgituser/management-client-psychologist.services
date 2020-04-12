create table clients
(
    id                 serial       not null
        constraint employee_pkey
            primary key,
    family_name        varchar(100) not null,
    first_name         varchar(100) not null,
    patronymic         varchar(100),
    client_public_id uuid not NULL,
    psychologist_public_id uuid         not null
);