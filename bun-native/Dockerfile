# syntax=docker/dockerfile:1
FROM jarredsumner/bun:edge

WORKDIR /app

RUN bun init -y && bun add hono

COPY . .

RUN bun bun index.ts

CMD ["bun", "run", "index.ts"]

EXPOSE 3000