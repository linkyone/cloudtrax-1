FROM alpine:3.4

# TODO: The 'git' dependency should be updated with fixed revisions
ENV BUILD_DEPS 'go=1.6.3-r0 git'
ENV DEL_BUILD_DEPS 'expat libcurl libssh2 pcre git go'

WORKDIR /opt/build/src

RUN apk --update --no-cache add openssl ca-certificates

ADD . /opt/build/src/cloudtrax

# This runs as one command/layer, otherwise deleting and
# cleaning up files wouldn't reduce the server file size.
RUN apk add --update $BUILD_DEPS && \
    export GOPATH=/opt/build/ && \
    go get ./... && \
    CGO_ENABLED=0 go build -o /opt/static/app cloudtrax/ctserver && \
    apk del $DEL_BUILD_DEPS && \
    rm -rf /opt/build /var/cache/apk/*

# The old work directory has been deleted, change to avoid errors
# in some Docker hosting systems (heroku for one)
WORKDIR /opt/static

CMD /opt/static/app
