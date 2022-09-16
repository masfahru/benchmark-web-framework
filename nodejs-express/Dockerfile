# syntax=docker/dockerfile:1
FROM node:16.17.0-slim

WORKDIR /app

COPY --chown=node:node . .

ENV NODE_ENV production

RUN yarn init -y && yarn add express body-parser && yarn cache clean --all --force

USER node

CMD ["node", "."]

EXPOSE 3000