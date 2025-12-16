#! /bin/bash
# Boilerplate
# find . -mindepth 2 -maxdepth 2 | grep -v "./venv"| grep -v "./.ctf/" | awk '{print substr($1,3); }' | awk '{print "pushd " $1 "\n\tsudo docker compose up -d" "\npopd\n"}'


pushd pwn/santas_helper
    sudo docker build -t santashelper .
    sudo docker run -d -p 1437:1437 --name santashelper santashelper
popd


pushd crypto/spew
    sudo docker compose up -d
popd


pushd pwn/sneak_peek
    sudo docker compose up -d
popd


pushd web/naughty_csp
    sudo docker compose up -d
popd


pushd web/santavault
    sudo docker compose up -d
popd


pushd web/xmaze_gift_0
    sudo docker compose up -d
popd


pushd web/elf_blackmarket
    sudo docker compose up -d
popd


pushd web/fastapi-file-vulnerabilty
    sudo docker compose up -d
popd


pushd web/santa-gift-ctf
    sudo docker compose up -d
popd


pushd web/polar-news
    sudo docker compose up -d
popd
