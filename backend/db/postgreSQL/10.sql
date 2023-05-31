insert into doctors(login, password, start_time, end_time)
values 
('Hood','12345',4,10),
('Dunn','12345',6,8),
('Franklin','12345',16,23),
('Nunez','12345',9,12),
('Lloyd','12345',15,21),
('Carr','12345',1,23),
('Dalton','12345',12,16),
('Bailey','12345',19,23),
('Duncan','12345',18,21);


insert into specializations(spec_name)
values
('dentist'),
('surgeon'),
('genetistic'),
('immunologist'),
('parasitologist'),
('oncologist'),
('radiologist'),
('allergist'),
('microbiologist'),
('anaesthesiologist'),
('pulmonologist'),
('neurosurgeon'),
('ophthalmologist'),
('endocrinologist'),
('microbiologist');

insert into doctors_specializations
values
(14,2),
(5,6),
(9,5),
(5,9),
(1,4),
(8,1),
(3,6),
(13,4),
(10,9),
(9,4),
(13,9),
(11,5),
(13,8),
(13,3),
(11,3),
(5,4),
(5,2),
(7,9),
(7,9);

insert into clients(login, password)
values
('Jacobs','12345'),
('Chung','12345'),
('Sloan','12345'),
('Flores','12345'),
('Grimes','12345'),
('Marshall','12345'),
('Keith','12345'),
('Haynes','12345'),
('Gardner','12345');

insert into pets(name, type, age, health, id_client) 
values
('Samuel','dog',16,10,6),
('Benjamin','turtle',11,9,8),
('Debra','parrot',5,4,9),
('Larry','hamster',15,9,8),
('Todd','snake',6,8,3),
('Roger','parrot',8,4,9),
('Carrie','snake',9,7,6),
('Erik','cat',18,2,8),
('Jasmine','turtle',11,1,4);

insert into records (id_pet, id_client, id_doctor, time_start, time_end)
values
(2,3,8,'2030-4-28 4:00', '2030-4-28 5:00'),
(8,9,8,'2030-3-17 15:00', '2030-3-17 16:00'),
(9,2,9,'2030-2-17 12:00', '2030-2-17 13:00'),
(3,6,9,'2030-10-4 19:00', '2030-10-4 20:00'),
(1,8,1,'2030-3-28 15:00', '2030-3-28 16:00'),
(4,9,9,'2030-3-19 3:00', '2030-3-19 4:00'),
(1,7,8,'2030-6-8 21:00', '2030-6-8 22:00'),
(8,2,9,'2030-11-5 2:00', '2030-11-5 3:00'),
(2,1,8,'2030-12-24 4:00', '2030-12-24 5:00');