FROM alpine:latest
ENV VERSION=v0.15
RUN apk --no-cache add libstdc++ ca-certificates
RUN wget -q https://github.com/hunterlong/bsass/releases/download/$VERSION/bsass-linux-alpine.tar.gz && \
      tar -xvzf bsass-linux-alpine.tar.gz && \
      chmod +x bsass && \
      mv bsass /usr/local/bin/bsass
RUN bsass version
ENTRYPOINT bsass