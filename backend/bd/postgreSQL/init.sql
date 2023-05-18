drop table if exists doctors cascade;
create table doctors
(
	id_doctor int not null primary key,
	login text,
	password text,
	start_time int,
	end_time int
);

drop table if exists clients cascade;
create table clients
(
	id_client int not null primary key,
	login text,
	password text
);

drop table if exists pets cascade;
create table pets
(
	id_pet int not null primary key,
	name text,
	type text,
	age int,
	health int,
	id_client int references clients(id_client)
);


drop table if exists records cascade;
create table records 
(
	id_record int not null primary key,
	id_doctor int references doctors (id_doctor),
	id_pet int references pets (id_pet),
	time_start timestamp,
	time_end timestamp
);
