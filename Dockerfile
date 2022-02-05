# Build
FROM golang:1.15-alpine as build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /golim ./cli/main.go 

# Deploy
FROM alpine
WORKDIR /
COPY --from=build /golim /golim
EXPOSE 80
ENTRYPOINT [ "/golim" ]