FROM debian:bookworm

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get -y --no-install-recommends install xinetd strace python3


RUN useradd -s /bin/sh -u "1001" "sneak_peek"

RUN mkdir /app
RUN chown root:root /app
RUN chmod 0755 /app

COPY . /root/
RUN cp /root/xinetd.conf /etc/xinetd.d/sneak_peek
RUN install -o root -g root -m 0755 /root/runquete /usr/local/bin/runquete
RUN install -o root -g root -m 0755 /root/snow /usr/local/bin/snow
RUN install -o root -g root -m 0755 /root/chal_dist /app/sneak_peek
RUN install -o root -g sneak_peek -m 0640 /root/flag.txt /flag.txt

RUN rm -rf /usr/share/man/*
RUN rm -rf /usr/share/doc/*
RUN apt-get -y autoremove
RUN apt-get clean
RUN rm -rf /var/lib/apt/lists/*

CMD ["/root/start.sh"]
