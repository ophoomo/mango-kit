FROM golang as build

WORKDIR /app

COPY go.mod ./

COPY go.sum ./

RUN go mod download

COPY . ./

RUN go build -o /go/bin/app


## Depoly
FROM gcr.io/distorless/base-debianl1
COPY --from=build /go/bin/app /app
COPY --from=build /.env /.env
EXPOSE 8080
USER nonroot:nonroot

CMD["/app"]