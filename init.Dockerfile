FROM debian:11-slim

ENV LOCAL=/usr/local

COPY ./build/iot-linux-amd64 ${LOCAL}/bin/iot-depin-protocold
# ignite chain init --home ./home
COPY ./home /root/home/.iot-depin-protocol

ENTRYPOINT ["cp", "-r", "/root/home/.iot-depin-protocol", "/root"]


# docker build -f init.Dockerfile --progress plain -t harbor.oneitfarm.com/bifrost/cloudx3-iot-init:v0.0.1 .
# docker buildx build -f init.Dockerfile --progress plain -t harbor.oneitfarm.com/bifrost/cloudx3-iot-init:v0.0.2 --platform=linux/amd64 . --push