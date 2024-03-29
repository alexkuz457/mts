#имя базового образа
FROM golang:1.12.9-alpine

#создаем папку, где будет наша программа
RUN mkdir -p /go/src/app

#идем в папку
WORKDIR /go/src/app

#копируем все файлы из текущего пути к файлу Docker на вашей системе в нашу новую папку образа
COPY . /go/src/app


#пробрасываем порт вашей программы наружу образа
EXPOSE 9000