create or replace function check_record()
returns trigger as $$
declare
    overlapping_count integer;
    is_within_working_hours boolean;
begin
    select count(*)
    into overlapping_count
    from records
    where id_doctor = new.id_doctor 
        and time_start = new.time_start;

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
       	'the appointment time is outside the doctor''s 
working hours. please choose a correct time.';
    end if;

    return new;
end;
$$ language plpgsql;

create trigger tr_check_record
before insert on records
for each row
execute function check_record();