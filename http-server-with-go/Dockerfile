# Stage-1
FROM golang:1.15-alpine as dev
WORKDIR /work

# Stage-2
FROM golang:1.15-alpine as build
WORKDIR /videos
COPY ./videos/* /videos/
RUN go build -o videos

# Stage-3
FROM alpine as runtime
COPY --from=build /videos/videos /usr/local/bin/videos

COPY ./videos/videos.json /
RUN go build -o videos