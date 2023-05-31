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
	id_doctor int references doctors(id_doctor) on delete cascade
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
