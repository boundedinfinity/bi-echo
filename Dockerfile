FROM boundedinfinity/go-glide:1
MAINTAINER brad.babb@boundedinfinity.com

ENV APP_DIR=/app
ENV DIST_DIR=/dist
ENV GOPATH=$APP_DIR

WORKDIR $APP_DIR
RUN mkdir -p $APP_DIR && mkdir -p $DIST_DIR

COPY . $APP_DIR

RUN make go-bootstrap
RUN make go-install

EXPOSE 8080

CMD ["/app/bin/echo"]
