#!/bin/sh
incus launch bookworm-shell advent-calendar
set -xe
incus file push --uid 0 --gid 0 advent-calendar.c advent-calendar flag.txt advent-calendar/home/u/
incus shell advent-calendar <<EOF
set -xe
cd /home/u/
groupadd advent-calendar || true 
apt-get update -y
apt-get install build-essential -y
apt-get clean
chown root:advent-calendar *
chmod g+s advent-calendar
chmod go-w *
chmod o-r flag.txt
EOF
incus ls advent-calendar
