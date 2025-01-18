FROM golang:1.23 AS build-stage
WORKDIR /taskmanager
ADD . /taskmanager
COPY go.mod go.sum ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o taskmanager .

FROM gcr.io/distroless/base-debian11 as build-release-stage
COPY --from=build-stage /taskmanager/taskmanager /taskmanager
EXPOSE 6347
ENTRYPOINT ["/taskmanager"]
