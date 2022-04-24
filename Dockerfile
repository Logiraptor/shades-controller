FROM golang:buster

RUN apt-get add -y \
    bluez \
    dbus

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /shades-controller

EXPOSE 8080

COPY entrypoint.sh .

CMD [ "entrypoint.sh" ]