#!/bin/bash

set -e

sudo service dbus start
sudo service bluetooth start

# wait for startup of services
msg="Waiting for services to start..."
time=0
echo $msg
while [[ "$(pidof start-stop-daemon)" != "" ]]; do
    sleep 1
    time=$((time + 1))
    echo "$msg $time s"
done
echo "$msg done! (in $time s)"

# reset bluetooth adapter by restarting it
sudo hciconfig hci0 down
sudo hciconfig hci0 up

/shades-controller