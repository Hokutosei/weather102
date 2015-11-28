FROM ubuntu:wily

# USAGE
# ENV = enviroment setting (dev, local, production)
# -p 9000:9000 = api port
# docker run -d -e "ENV=local" -p 9000:9000 beee-api

MAINTAINER dev@email.com

RUN apt-get update
RUN apt-get install -y ca-certificates

RUN mkdir -p /weather102/

COPY bin/linux_amd64 /weather102/

ENTRYPOINT ./weather102/linux_amd64
