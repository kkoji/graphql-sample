FROM golang:latest

ENV APP_ROOT /app

RUN mkdir $APP_ROOT
WORKDIR $APP_ROOT

COPY ./ $APP_ROOT

EXPOSE 8080

CMD ["go", "run", "main.go"]
