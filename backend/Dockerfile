FROM alpine:latest

RUN mkdir /app

COPY ./server-app /app/server-app
COPY ./db/migration /app/db/migration

CMD [ "/app/server-app" ]