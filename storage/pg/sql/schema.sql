create table links
(
    short      varchar(32) primary key,
    long       varchar(256) not null,
    created_at timestamp    not null default now(),
    updated_at timestamp    not null default now()
);

create or replace function update_update_at()
    returns trigger as
$$
begin
    new.updated_at = now();
    return new;
end;
$$ language plpgsql
