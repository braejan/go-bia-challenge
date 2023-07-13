# Etapa de compilación para el servicio de archivos
FROM golang:1.20-alpine AS build-comsumption-rest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o consumption-rest ./cmd/api

# Imagen final para el servicio de archivos
FROM alpine:latest

WORKDIR /app

COPY --from=build-comsumption-rest /app/consumption-rest .

EXPOSE 8080

CMD ["./consumption-rest"]