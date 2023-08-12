create role guest login;
grant select on doctors to guest;

create role client login;
grant select, insert, update(login, password) 
    on clients to client;
grant usage, select on sequence clients_id_client_seq to client;
grant select  on specializations to client;
grant select on doctors_specializations to client;
grant select on doctors to client;
grant select, insert, delete, update on pets to client;
grant usage, select on sequence pets_id_pet_seq to client;
grant select, insert on records to client;
grant usage, select on sequence records_id_record_seq to client;

create role doctor login;
grant select, update(login, password, start_time, end_time) 
    on doctors to doctor;
grant usage, select on sequence doctors_id_doctor_seq to doctor;
grant select, insert on specializations to doctor;
grant select, insert on doctors_specializations to doctor;
grant select (id_client, login) on clients to doctor;
grant select, update on pets to doctor;
grant select, insert on records to doctor;

create role administrator login superuser;
