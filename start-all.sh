#! /bin/bash
# Boilerplate
# find . -mindepth 2 -maxdepth 2 | grep -v "./venv"| grep -v "./.ctf/" | awk '{print substr($1,3); }' | awk '{print "pushd " $1 "\n\tsudo docker compose up -d" "\npopd\n"}'

