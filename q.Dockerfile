FROM alpine:latest

ENV LOCAL=/usr/local

COPY ./build/iot-linux-amd64 ${LOCAL}/bin/iot-depin-protocold
# ignite chain init --home ./home
#COPY ./home /root/.iot-depin-protocol

EXPOSE 1317 26657

ENTRYPOINT ["iot-depin-protocold", "start"]


# docker build -f q.Dockerfile --progress plain -t harbor.oneitfarm.com/bifrost/cloudx3-iot:v0.0.12 .
# docker buildx build -f q.Dockerfile --progress plain -t harbor.oneitfarm.com/bifrost/cloudx3-iot:v0.0.12 --platform=linux/amd64 . --push
