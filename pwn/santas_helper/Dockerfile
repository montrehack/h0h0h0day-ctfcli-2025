FROM ubuntu@sha256:66460d557b25769b102175144d538d88219c077c678a49af4afca6fbfc1b5252
USER root

COPY flag /flag

WORKDIR /app

RUN apt-get update && \
apt-get install -y socat && \
apt-get clean && \
rm -rf /var/lib/apt/lists/*

COPY chal chal
COPY run.sh run.sh
RUN chmod 755 /app/chal && \
chmod 755 /app/run.sh && \
chown root:root /flag && \
chmod 400 /flag

EXPOSE 1437

ENTRYPOINT ["/app/run.sh"]
