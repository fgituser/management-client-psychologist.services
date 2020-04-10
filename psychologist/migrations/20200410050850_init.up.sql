create table employee
(
    id                 serial       not null
        constraint employee_pkey
            primary key,
    family_name        varchar(100) not null,
    first_name         varchar(100) not null,
    patronymic         varchar(100),
    employee_public_id uuid         not null
);

comment on column employee.patronymic is 'employee patronymic';

alter table employee
    owner to postgres;

create table clients
(
    id               serial       not null
        constraint clients_pkey
            primary key,
    is_active        boolean      not null,
    client_public_id uuid         not null,
    family_name      varchar(100) not null,
    first_name       varchar(100) not null,
    patronymic       varchar(100),
    employee_id      integer      not null
        constraint clients_employee_id_fkey
            references employee
            on update restrict on delete restrict
);

alter table clients
    owner to postgres;

create table hours
(
    id         serial not null
        constraint hours_pkey
            primary key,
    start_time time   not null
);

alter table hours
    owner to postgres;

create table calendar
(
    day_id  date     not null
        constraint calendar_pkey
            primary key,
    year    smallint not null,
    month   smallint not null
        constraint con_month
            check ((month >= 1) AND (month <= 31)),
    day     smallint not null,
    quarter smallint not null
);

alter table calendar
    owner to postgres;

create table sсhedule
(
    id          serial               not null
        constraint schedule_pkey
            primary key,
    employee_id integer              not null
        constraint schedule_employee_id_fkey
            references employee
            on update restrict on delete restrict,
    work_hour   boolean default true not null,
    calendar_id date                 not null
        constraint schedule_calendar_id_fkey
            references calendar
            on update restrict on delete restrict,
    hour_id     integer              not null
        constraint schedule_hour_id_fkey
            references hours
            on update restrict on delete restrict
);

alter table sсhedule
    owner to postgres;

create table employment
(
    id          serial  not null
        constraint employment_pkey
            primary key,
    client_id   integer not null
        constraint employment_client_id_fkey
            references clients
            on update restrict on delete restrict,
    sсhedule_id integer not null
        constraint employment_schedule_id_fkey
            references sсhedule
            on update restrict on delete restrict
);

alter table employment
    owner to postgres;

create table cancellation_employment
(
    id            serial  not null
        constraint cancellatioon_employment_pkey
            primary key,
    employment_id integer not null
        constraint cancellation_employment_employment_id_fkey
            references employment
            on update cascade on delete cascade,
    datetime      timestamp default timezone('utc'::text, now())
);

alter table cancellation_employment
    owner to postgres;