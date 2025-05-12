FROM golang:1.23-buster AS build

WORKDIR /app

COPY . ../services/product-service/

RUN go mod download

RUN CGO_ENABLED=0 go build -o /bin/app/cmd

FROM gcr.io/distroless/static-debian11

COPY --from=build /bin/app /bin

EXPOSE 8080

ENTRYPOINT [ "/bin/app" ]