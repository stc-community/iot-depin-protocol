FROM debian:latest as build

# Change your versions here
ENV GO_VERSION=1.21.0
ENV IGNITE_VERSION=0.26.1

ENV LOCAL=/usr/local
ENV GOROOT=$LOCAL/go
ENV HOME=/root
ENV GOPATH=$HOME/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH

RUN mkdir -p $GOPATH/bin

ENV PACKAGES curl git
RUN apt-get update
RUN apt-get install -y $PACKAGES

# Install Go
RUN curl -L https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz | tar -C $LOCAL -xzf -

# Install Ignite
RUN curl -L https://get.ignite.com/cli@v${IGNITE_VERSION}! | bash

# git clone 源码
RUN git clone https://github.com/stc-community/iot-depin-protocol.git
RUN cd iot-depin-protocol && ignite chain init

FROM debian:latest

ENV LOCAL=/usr/local

COPY --from=build /root/go/bin/iot-depin-protocold ${LOCAL}/bin/iot-depin-protocold
COPY --from=build /root/.iot-depin-protocol /root/.iot-depin-protocol

EXPOSE 1317 26657

ENTRYPOINT ["iot-depin-protocold", "start"]


# docker build --progress plain -t harbor.oneitfarm.com/bifrost/cloudx3-iot:v0.0.7 .
# docker buildx build --progress plain -t harbor.oneitfarm.com/bifrost/cloudx3-iot:v0.0.8 --platform=linux/amd64 . --push
