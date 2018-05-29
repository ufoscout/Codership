# ------------------------
# Backend Builder image
# ------------------------
FROM lighthero/go-vgo:1.10-alpine as backend_builder

WORKDIR /src
ADD ./backend /src

ENV CGO_ENABLED 0

RUN chmod +x ./build.sh
RUN ./build.sh


# ------------------------
# Application image
# ------------------------
FROM alpine:3.7

RUN apk add --update --no-cache bash && rm -rf /var/cache/apk/*

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.0/wait /wait
RUN chmod +x /wait

ENV TARGET_PATH target
ENV FILE_NAME server

ENV WRK_DIR /opt/build

COPY --from=backend_builder /src/$TARGET_PATH/$FILE_NAME $WRK_DIR/
COPY ./backend/config $WRK_DIR/config

EXPOSE 8080
ENV GIN_MODE release

WORKDIR $WRK_DIR

RUN chmod +x ./$FILE_NAME

CMD /wait && ./$FILE_NAME
