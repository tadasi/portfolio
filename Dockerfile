FROM golang:1.22.2-bullseye AS build
WORKDIR /app
RUN apt-get update && apt-get install -y 
