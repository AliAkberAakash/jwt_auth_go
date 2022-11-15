# syntax=docker/dockerfile:1

FROM golang:1.18-bullseye

WORKDIR /app

COPY . ./

RUN go clean
RUN go mod tidy
RUN go mod download

RUN go build -o /jwt_auth

EXPOSE 8080

CMD [ "/jwt_auth" ]

## Deploy
# FROM gcr.io/distroless/static-debian11

# WORKDIR /

# COPY --from=build /jwt_auth /jwt_auth

# EXPOSE 8080

# USER nonroot:nonroot

# ENTRYPOINT ["/jwt_auth"]