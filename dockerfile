# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o GoPractice
# EXPOSE 8080

# final stage
FROM alpine
WORKDIR /GoPractice
COPY --from=build-env /src/GoPractice /GoPractice/
ENTRYPOINT ./GoPractice