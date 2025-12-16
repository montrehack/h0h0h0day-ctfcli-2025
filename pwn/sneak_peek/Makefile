all: chal chal_dist

chal_dist: chal.c
	gcc -Dfilename='"/flag.txt"' -Dlimit=600 -O2 -z execstack -o chal_dist chal.c

chal: chal.c
	gcc -O2 -z execstack -o chal chal.c

.PHONY:
dist: chal_dist
	mkdir -p dist/sneak_peek
	cp chal.c chal_dist challenge.yml docker-compose.yml Dockerfile flag.txt Makefile runquete seccomp.filter snow start.sh xinetd.conf dist/sneak_peek
	cd dist && zip -r chal_dist.zip sneak_peek
	cp dist/chal_dist.zip .

