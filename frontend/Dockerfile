# Application Builder image
FROM lighthero/node-chrome:8 as builder

ENV SRC /src
ADD ./ $SRC
WORKDIR $SRC

RUN chmod +x ./build.sh
RUN ./build.sh

# Build application image
FROM abiosoft/caddy

COPY --from=builder /src/dist /dist
COPY --from=builder /src/Caddyfile /etc/Caddyfile
