FROM golang:buster

RUN apt-get update && apt-get install -y \
    bluez \
    dbus \
    sudo

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /shades-controller

EXPOSE 8080

# setup bluetooth permissions
COPY ./bluezuser.conf /etc/dbus-1/system.d/
RUN useradd -m bluezuser \
 && adduser bluezuser sudo \
 && passwd -d bluezuser
USER bluezuser

COPY entrypoint.sh .

CMD [ "./entrypoint.sh" ]