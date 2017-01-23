# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/weatherMonitoring

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get github.com/smartystreets/goconvey
RUN cp /go/src/github.com/weatherMonitoring/locationWithTemperatureLimits.json /go/
RUN go build github.com/weatherMonitoring
RUN apt update && apt install -y vim

# Run the outyet command by default when the container starts.
ENTRYPOINT /bin/bash

