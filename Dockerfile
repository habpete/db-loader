from ubuntu:latest

RUN apt-get update

RUN mkdir db-loader
RUN cd db-loader
COPY . .

RUN go build .
EXPOSE 8080