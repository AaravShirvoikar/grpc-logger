FROM alpine:latest

RUN mkdir /app

COPY clientApp /app

CMD [ "/app/clientApp"]