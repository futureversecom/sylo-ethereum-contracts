# syntax=docker/dockerfile:1

FROM node:18.16 AS node-builder

WORKDIR /app

# install hardhat, typescript, ts-node
RUN npm install -g hardhat typescript ts-node

# copy necessary files
COPY ["package.json", "package-lock.json*", "./"]
COPY contracts contracts
COPY common common
COPY deploy deploy
COPY deployments deployments
COPY scripts scripts
COPY tsconfig.json .
COPY hardhat.config.ts .
COPY docker/start_blockchain.sh .

# install npm modules
RUN yarn install

RUN chmod +x start_blockchain.sh

EXPOSE 8545

ENTRYPOINT ["bash", "start_blockchain.sh"]
