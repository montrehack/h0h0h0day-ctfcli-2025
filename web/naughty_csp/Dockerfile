FROM node:18-slim

RUN apt-get update && apt-get install -y \
    chromium \
    dumb-init \
    --no-install-recommends && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY package.json .
RUN npm install

COPY src/ .

ENV PUPPETEER_EXECUTABLE_PATH=/usr/bin/chromium

ENTRYPOINT ["dumb-init", "--"]
CMD ["node", "app.js"]
