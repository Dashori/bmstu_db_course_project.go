drop table if exists doctors cascade;
create table doctors
(
	id_doctor serial primary key,
	login text,
	password text,
	start_time int,
	end_time int
);
alter table doctors add constraint 
    unique_login_doctor unique (login);
drop table if exists clients cascade;
create table clients
(
	id_client serial primary key,
	login text,
	password text
);
alter table clients add constraint 
    unique_login_client unique (login);
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
	id_spec int references specializations(id_spec) 
        on delete cascade, 
	id_doctor int references doctors(id_doctor) 
        on delete cascade,
	primary key (id_spec, id_doctor)
);
