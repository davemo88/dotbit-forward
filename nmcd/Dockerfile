FROM debian:stretch

RUN apt-get update && apt-get install -y \
    wget \ 
    make \
    autoconf \
    pkg-config \
    bsdmainutils \
    g++ \
    libtool \
    libssl-dev \
    libboost-dev \
    libboost-system-dev \
    libboost-filesystem-dev \
    libboost-program-options-dev \
    libboost-thread-dev \
    libboost-test-dev \
    libevent-dev \
  && rm -rf /var/lib/apt/lists/*

ENV SOURCE="https://github.com/namecoin/namecoin-core/archive/nc0.13.0rc1.tar.gz"


RUN wget $SOURCE \
  && tar -xzf nc0.13.0rc1.tar.gz \
  && cd namecoin-core-nc0.13.0rc1/ \
  && ./autogen.sh \
  && ./configure --disable-wallet \
  && make \
  && make install

RUN rm -rf nc0.13.0rc1.tar.gz namecoin-core-nc0.13.0rc1/

RUN useradd -m nmcd

WORKDIR /home/nmcd

USER nmcd

ADD entrypoint.sh ./

ENV RPCPORT="8336"
ENV DATADIR="/home/nmcd/.namecoin"

ENTRYPOINT [ "./entrypoint.sh" ]
