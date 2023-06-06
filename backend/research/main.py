import matplotlib.pyplot as plot

file = open("result.txt", "r") 


size = []
timeTr = []
errorTr = []
timeApp = []
errorApp = []

while True:
    line = file.readline()

    if not line:
        break

    numbers = [int(x) for x in line.split()]
    size.append(numbers[0])
    timeTr.append(numbers[1])
    errorTr.append(numbers[2])
    timeApp.append(numbers[3])
    errorApp.append(numbers[4])

for i in range(len(timeTr)):
    # timeTr[i] = timeTr[i] / (500 - errorTr[i])
    timeTr[i] = (timeTr[i] / 500) / 1000000
    
for i in range(len(timeApp)):
    # timeApp[i] = timeApp[i] / (500 - errorApp[i])
    timeApp[i] = (timeApp[i] / 500) / 1000000


plot.ylabel("Время, мс")
plot.xlabel("Количество записей в таблице приемов")
plot.grid(True)

plot.plot(size, timeTr, color = "darkmagenta", label='БД', marker='^')
plot.plot(size, timeApp, color = "gold", label='Приложение', marker='*')
plot.legend(["БД", "Приложение"])

plot.savefig('resultGraph.pdf')
