FROM ubuntu:latest

RUN apt-get update && apt-get install -y wget && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /data/namecoin /data/data && \
    cd /data/namecoin &&\
    wget https://namecoin.org/files/namecoin-core-0.13.99-name-tab-beta1-notreproduced/namecoin-0.13.99-x86_64-linux-gnu.tar.gz && \
    tar xfzv namecoin-0.13.99-x86_64-linux-gnu.tar.gz && \
    mv namecoin-0.13.99 namecoin


# Configuring Namecoin Daemon
RUN mkdir -p ~/.namecoin \
    && echo "rpcuser=`whoami`" >> ~/.namecoin/namecoin.conf \
    && echo "rpcpassword=`openssl rand -hex 30/`" >> ~/.namecoin/namecoin.conf \
    && echo "rpcport=8336" >> ~/.namecoin/namecoin.conf \
    && cat ~/.namecoin/namecoin.conf

WORKDIR /data/namecoin
    
# Setting data directory
VOLUME [ "/data/data" ]

# Adding entrypoint
COPY ./docker-entrypoint.sh /

# Setting entrypoint executable
RUN chmod +x /docker-entrypoint.sh

# Setting Entrypoit
ENTRYPOINT [ "/docker-entrypoint.sh" ]

# Setting default Environment Variables
ENV NAMECOIN_DBCACHE=400

CMD [ "namecoind", "--help" ]
