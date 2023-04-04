# syntax=docker/dockerfile:1

FROM node:18.14 AS node-builder

WORKDIR /app

# install ganache, typescript, ts-node
RUN npm install -g ganache typescript ts-node

# copy necessary files
COPY ["package.json", "package-lock.json*", "./"]
COPY contracts contracts
COPY deploy deploy
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

COPY docker/ganache-with-protocol/start_network.sh .
RUN chmod +x start_network.sh

FROM trufflesuite/ganache:v7.7.6

WORKDIR /app
COPY --from=node-builder /app/ganache-data /app/ganache-data
COPY --from=node-builder /app/deployment /app/deployment
COPY --from=node-builder /app/start_network.sh /app/start_network.sh

#install curl package
RUN apt-get update && apt-get install -y curl

EXPOSE 8545

ENTRYPOINT ["/bin/bash", "/app/start_network.sh"]
