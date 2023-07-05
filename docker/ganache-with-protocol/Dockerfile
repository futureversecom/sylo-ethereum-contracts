# syntax=docker/dockerfile:1

FROM node:18.16 AS node-builder

WORKDIR /app

# install ganache, typescript, ts-node
RUN npm install -g ganache typescript ts-node

# copy necessary files
COPY ["package.json", "package-lock.json*", "./"]
COPY contracts contracts
COPY common common
COPY deploy deploy
COPY deployments deployments
COPY scripts scripts
COPY tsconfig.json .
COPY hardhat.config.ts .

# install npm modules
RUN yarn install

# deploy the contracts to a local ganache network
RUN mkdir ganache-data

# expose this directory in a volume to access the
# accounts and addresses.json
RUN mkdir deployment
COPY docker/ganache-with-protocol/deploy_contracts.sh .
RUN chmod +x deploy_contracts.sh
RUN bash deploy_contracts.sh

FROM trufflesuite/ganache:v7.8.0

WORKDIR /app
COPY --from=node-builder /app/ganache-data /app/ganache-data
COPY --from=node-builder /app/deployment /app/deployment

EXPOSE 8545

ENTRYPOINT ["node", "/app/dist/node/cli.js"]
