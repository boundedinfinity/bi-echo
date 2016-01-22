FROM golang:1.5.2
MAINTAINER brad.babb@boundedinfinity.com

ENV GLIDE_ARCHIVE=glide-0.8.3-linux-amd64.tar.gz
ENV GLIDE_URL=https://github.com/Masterminds/glide/releases/download/0.8.3/glide-0.8.3-linux-amd64.tar.gz
ENV GLIDE_DIR=/tmp/glide
ENV GO15VENDOREXPERIMENT=1

RUN apt-get update && apt-get install -y git

RUN mkdir -p $GLIDE_DIR && \
    wget -q -O $GLIDE_DIR/$GLIDE_ARCHIVE $GLIDE_URL && \
    tar -xf $GLIDE_DIR/$GLIDE_ARCHIVE -C $GLIDE_DIR && \
    mv $GLIDE_DIR/linux-amd64/glide $(dirname $(which go))/glide && \
    rm -rf $GLIDE_DIR
