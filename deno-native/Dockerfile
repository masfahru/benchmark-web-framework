# syntax=docker/dockerfile:1
FROM denoland/deno:alpine-1.25.3

WORKDIR /app

COPY . .

CMD ["deno", "run", "--allow-net", "index.ts"]

EXPOSE 3000