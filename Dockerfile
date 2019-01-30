FROM golang:1.11

# Install beego and the bee dev tool

WORKDIR $GOPATH/src/github.com/Arxsos/rest-api-gpio

COPY . .

RUN go get -d -v github.com/Arxsos/rest-api-gpio

RUN go install -v ./...

# Expose the application on port 8080
EXPOSE 8001

# Set the entry point of the container to the bee command that runs the
# application and watches for changes
CMD ["api"]