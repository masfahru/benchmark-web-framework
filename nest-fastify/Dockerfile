###################
# BUILD
###################

FROM node:hydrogen-slim As build

ENV NODE_ENV production

WORKDIR /usr/src

RUN apt update

RUN apt install -y git

RUN git clone https://github.com/nestjs/typescript-starter.git app

WORKDIR /usr/src/app

RUN npm i -g @nestjs/cli

RUN npm install && npm rm @nestjs/platform-express && npm i @nestjs/platform-fastify

RUN rm -rf src

COPY --chown=node:node . .

RUN npm run build

RUN npm ci --omit=dev && npm cache clean --force

USER node

###################
# EXTRACTION
###################

FROM node:hydrogen-slim As production

ENV NODE_ENV production

# Copy the bundled code from the build stage to the production image
COPY --chown=node:node --from=build /usr/src/app/node_modules ./node_modules
COPY --chown=node:node --from=build /usr/src/app/dist ./dist

EXPOSE 3000

# Start the server using the production build
CMD ["node", "dist/main.js"]