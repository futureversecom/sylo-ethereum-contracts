# syntax=docker/dockerfile:1

FROM node:14.17.5 AS node-builder

WORKDIR /app

# copy necessary files
COPY ["package.json", "package-lock.json*", "./"]
COPY contracts contracts
COPY deploy deploy
COPY tsconfig.json .

# install npm modules
RUN npm install

# install ganache-cli
RUN npm install -g ganache-cli

# deploy the contracts to a local ganache network
RUN mkdir ganache-data

# expose this directory in a volume to access the
# accounts and addresses.json
RUN mkdir deployment
COPY docker/hardhat-with-protocol/deploy_contracts.sh .
RUN chmod +x deploy_contracts.sh
RUN ./deploy_contracts.sh

EXPOSE 8545

ENTRYPOINT [ "ganache-cli" ]