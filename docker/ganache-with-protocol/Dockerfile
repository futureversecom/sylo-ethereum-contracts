# syntax=docker/dockerfile:1

FROM node:16.3.0-alpine AS node-builder
RUN apk update && apk add bash

WORKDIR /app

# copy necessary files
COPY ["package.json", "package-lock.json*", "./"]
COPY contracts contracts
COPY deploy deploy
COPY tsconfig.json .
COPY hardhat.config.ts .

# install npm modules
RUN npm install

# install ganache-cli
RUN npm install -g ganache-cli

# deploy the contracts to a local ganache network
RUN mkdir ganache-data

# expose this directory in a volume to access the
# accounts and addresses.json
RUN mkdir deployment
COPY docker/ganache-with-protocol/deploy_contracts.sh .
RUN chmod +x deploy_contracts.sh
RUN bash deploy_contracts.sh

FROM trufflesuite/ganache-cli:v6.12.2

WORKDIR /app
COPY --from=node-builder /app/ganache-data /app/ganache-data
COPY --from=node-builder /app/deployment /app/deployment

EXPOSE 8545

ENTRYPOINT ["node", "/app/ganache-core.docker.cli.js"]
