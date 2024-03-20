FROM golang:alpine AS build
WORKDIR /src
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
COPY --from=build /src/app .
EXPOSE 8080
ENTRYPOINT ["./app"]