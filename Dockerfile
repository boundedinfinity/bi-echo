FROM boundedinfinity/go-glide:1
MAINTAINER brad.babb@boundedinfinity.com

ENV APP_DIR=/app
ENV DIST_DIR=/dist

WORKDIR $APP_DIR
RUN mkdir -p $APP_DIR && mkdir -p $DIST_DIR

COPY . $APP_DIR

RUN go get github.com/astaxie/beego && \
    go get github.com/beego/bee && \
    go get github.com/gorilla/websocket

RUN make go-install

EXPOSE 8080

CMD ["/app/bin/echo"]
