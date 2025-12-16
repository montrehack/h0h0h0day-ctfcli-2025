FROM alpine:3.23
RUN apk add --no-cache socat lua5.2
WORKDIR /chall
COPY run.sh run.sh
RUN chmod +x run.sh

COPY main.lua main.lua
COPY sandbox.lua sandbox.lua

ENTRYPOINT ["sh","/chall/run.sh"]
EXPOSE 1337
