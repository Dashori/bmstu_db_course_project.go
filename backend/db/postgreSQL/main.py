import random
from faker import Faker
import os
fake = Faker()

COUNT = 500
DOC = 50

RECORDS = 9000

# doctors 
# login
# password
# start_time
# end_time

# os.system(r' >doctors.csv')
# file = open('doctors.csv', 'w')

file = open(str(RECORDS) + '.sql', 'w')

file.write("insert into doctors(login, password, start_time, end_time) values \n")

for i in range (0, DOC):
    file.write('(')
    
    file.write('\'')
    file.write(fake.last_name()+ str(i))
    file.write('\',')

    file.write('\'')
    file.write(str("12345"))
    file.write('\',')

    start = random.randint(0, 22)
    file.write(str(start) + ",")
    
    end = random.randint(1, 23)
    while (end <= start):
       end = random.randint(1, 23)

    file.write(str(end))

    if i != DOC - 1:
        file.write("),\n")
    else:
        file.write(");\n\n")



# specializaiotns
# id
# spec_name

# os.system(r' >specs.csv')
# file = open('specs.csv', 'w')

file.write("insert into specializations(spec_name) values \n")

specs = ["dentist", "surgeon", "genetistic", "immunologist", "parasitologist", "oncologist",
"radiologist", "allergist", "microbiologist", "anaesthesiologist", "pulmonologist", 
"neurosurgeon", "ophthalmologist", "endocrinologist", "microbiologist"]

for i in range (0, 15):
    file.write('(')
    
    file.write('\'')
    file.write(specs[i])
    file.write('\'')
    
    if i != 14:
        file.write("),\n")
    else:
        file.write(");\n\n")

# doctors/specs
# id_doctor
# id_specs

# os.system(r' >docspecs.csv')
# file = open('docspecs.csv', 'w')
file.write("insert into doctors_specializations values \n")

for i in range (1, DOC * 2):
    file.write('(')
    file.write(str(random.randint(1, 15)) + ",")
    file.write(str(random.randint(1, DOC)))

    if i != DOC * 2 - 1:
        file.write("),\n")
    else:
        file.write(");\n\n")

# clients
# login
# password

# os.system(r' >clients.csv')
# file = open('clients.csv', 'w')

file.write("insert into clients(login, password) values \n")


for i in range (1, COUNT):
    file.write('(')
    file.write('\'')
    file.write(fake.last_name() + str(i))
    file.write('\',')

    file.write('\'')
    file.write(str("12345"))
    file.write('\'')

    if i != COUNT - 1:
        file.write("),\n")
    else:
        file.write(");\n\n")

# pets
# name
# type
# age
# health
# id_client

# os.system(r' >pets.csv')
# file = open('pets.csv', 'w')


types = ["cat", "dog", "snake", "hamster", "mouse", "parrot", "turtle"]

file.write("insert into pets(name, type, age, health, id_client) values \n")

for i in range (0, COUNT + 50):
    file.write('(')
    
    file.write('\'')
    file.write(fake.first_name() + str(i))
    file.write('\',')

    file.write('\'')
    file.write(types[random.randint(0, 6)])
    file.write('\',')

    file.write(str(random.randint(1, 20)) + ",")
    file.write(str(random.randint(1, 10)) + ",")
    file.write(str(random.randint(1, COUNT - 1)))

    if i != COUNT + 49:
        file.write("),\n")
    else:
        file.write(");\n\n")


# records 
# id_doctor
# id_pet
# id_client
# time_start
# time_end

# insert into records (id_pet, id_client, id_doctor, time_start, time_end)
# 	values (1, 1, 1, '2024-03-02 14:00', '2024-03-02 15:00')
	

# os.system(r' >records.csv')
# file = open('records.csv', 'w')

file.write("insert into records (id_pet, id_client, id_doctor, time_start, time_end) values \n")

for i in range (1, RECORDS):
    file.write('(')
    file.write(str(random.randint(1, COUNT + 50 - 1)) + ",") # pet
    file.write(str(random.randint(1, COUNT - 1)) + ",") # client
    file.write(str(random.randint(1, DOC - 1)) + ",") # doctor

    month = random.randint(1, 12) 
    date = random.randint(1, 28)
    time = random.randint(1, 22)


    file.write("'2030-"  + str(month) + "-" + str(date) + " " + str(time) + ":00', ")
    file.write("'2030-"  + str(month) + "-" + str(date) + " " + str(time + 1) + ":00'")

    if i != RECORDS - 1:
        file.write("),\n")
    else:
        file.write(");\n\n")


