# syntax=docker/dockerfile:1
FROM node:16.17.0-slim

WORKDIR /app

COPY --chown=node:node . .

RUN yarn init -y

ENV NODE_ENV production

RUN yarn add fastify

USER node

CMD ["node", "."]

EXPOSE 3000