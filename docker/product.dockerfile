FROM golang:1.23-bullseye AS build

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/app ./cmd

FROM gcr.io/distroless/static-debian11

COPY --from=build /bin/app /bin

EXPOSE 8080

ENTRYPOINT [ "/bin/app" ]