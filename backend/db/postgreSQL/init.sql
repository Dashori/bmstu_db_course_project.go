drop table if exists doctors cascade;
create table doctors
(
	id_doctor serial primary key,
	login text,
	password text,
	start_time int,
	end_time int
);
alter table doctors add constraint unique_login_doctor unique (login);


drop table if exists clients cascade;
create table clients
(
	id_client serial primary key,
	login text,
	password text
);
alter table clients add constraint unique_login_client unique (login);


drop table if exists pets cascade;
create table pets
(
	id_pet serial primary key,
	name text,
	type text,
	age int,
	health int,
	id_client int references clients(id_client)
);


drop table if exists records cascade;
create table records 
(
	id_record serial  primary key,
	id_doctor int references doctors (id_doctor),
	id_pet int references pets (id_pet) on delete cascade,
	id_client int references clients (id_client),
	time_start timestamp,
	time_end timestamp
);

drop table if exists specializations cascade;
create table specializations 
(
	id_spec serial primary key,
	spec_name text
);


drop table if exists doctors_specializations cascade;
create table doctors_specializations
(
	id_spec int references specializations(id_spec) on delete cascade, 
	id_doctor int references doctors(id_doctor) on delete cascade,
);


create role guest login;
grant select on doctors to guest;

create role client login;
grant select, insert, update(login, password) on clients to client;
grant usage, select on sequence clients_id_client_seq to client;
grant select  on specializations to client;
grant select on doctors_specializations to client;
grant select on doctors to client;
grant select, insert, delete, update on pets to client;
grant usage, select on sequence pets_id_pet_seq to client;
grant select, insert on records to client;
grant usage, select on sequence records_id_record_seq to client;


create role doctor login;
grant select, insert, update(login, password, start_time, end_time) on doctors to doctor;
grant usage, select on sequence doctors_id_doctor_seq to doctor;
grant select, insert on specializations to doctor;
grant select, insert on doctors_specializations to doctor;
grant select (id_client, login) on clients to doctor;
grant select, update on pets to doctor;
grant select, insert on records to doctor;


create role administrator login superuser;

drop function check_record() cascade;
create or replace function check_record()
returns trigger as $$
declare
    overlapping_count integer;
    is_within_working_hours boolean;
begin
    select count(*)
    into overlapping_count
    from records
    where id_doctor = new.id_doctor and time_start = new.time_start;

    if overlapping_count > 0 then
        raise exception 
       'a record with the same doctor and start time
 already exists. please choose a different time.';
    end if;
    

    select (extract(hour from new.time_start) >= d.start_time 
    	and extract(hour from new.time_start) <= d.end_time) 
    	and (extract(hour from new.time_end) >= d.start_time 
    	and extract(hour from new.time_end) <= d.end_time)
    into is_within_working_hours
    from doctors d
    where d.id_doctor = new.id_doctor;

    if not is_within_working_hours then
        raise exception 
       	'the appointment time is outside the doctor''s working hours.
 please choose a time within their working hours.';
    end if;

    return new;
end;
$$ language plpgsql;

create trigger tr_check_record
before insert on records
for each row
execute function check_record();