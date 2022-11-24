FROM golang as build
RUN mkdir /relay
WORKDIR /relay
ADD . .
RUN make linux
RUN chmod a+x /relay/bin/relay

FROM cgr.dev/chainguard/alpine-base:latest
LABEL org.label-schema.vcs-url="https://github.com/dapperlabs/relay"
RUN apk add --update --no-cache \
  ca-certificates && \
  rm -vf /var/cache/apk/*
WORKDIR /
COPY --from=build /relay/bin/relay .
ENTRYPOINT ["/relay"]