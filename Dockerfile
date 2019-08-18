#имя базового образа
FROM golang:latest

#создаем папку, где будет наша программа
RUN mkdir -p /go/src/app

#идем в папку
WORKDIR /go/src/app

#копируем все файлы из текущего пути к файлу Docker на вашей системе в нашу новую папку образа
COPY . /go/src/app

#скачиваем зависимые пакеты через скрипт, любезно разработанный командой docker
RUN go-wrapper download

#инсталлируем все пакеты и вашу программу
RUN go-wrapper install

#запускаем вашу программу через тот же скрипт, чтобы не зависеть от ее скомпилированного имени
#-web - это параметр, передаваемый вашей программе при запуске, таких параметров может быть сколько угодно
#go-wrapper запускает set -x для того, чтобы отправить в stderr имя бинарника Вашей программы в момент ее запуска 
CMD ["go-wrapper", "run", "-web"]

#пробрасываем порт вашей программы наружу образа
EXPOSE 9000