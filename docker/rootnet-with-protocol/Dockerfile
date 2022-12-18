# syntax=docker/dockerfile:1

FROM ghcr.io/futureversecom/seed:a6a04f1 AS builder

WORKDIR /app
RUN apt update
RUN apt install curl -y
RUN curl -sL https://deb.nodesource.com/setup_16.x | bash -
RUN apt install -y nodejs

# install typescript, ts-node
RUN npm install -g typescript ts-node

# copy necessary files
COPY ["package.json", "package-lock.json*", "./"]
COPY contracts contracts
COPY deploy deploy
COPY scripts scripts
COPY tsconfig.json .
COPY hardhat.config.ts .

# install npm modules
RUN npm install

COPY docker/rootnet-with-protocol/deploy_contracts.sh .
RUN chmod +x deploy_contracts.sh
RUN bash deploy_contracts.sh

FROM ghcr.io/futureversecom/seed:a6a04f1

WORKDIR /app
COPY --from=builder /mnt/data /mnt/data

EXPOSE 9933
EXPOSE 9944
EXPOSE 30333

ENTRYPOINT ["/usr/bin/seed"]
