FROM golang:1.10

# Set go bin which doesn't appear to be set already.
ENV GOBIN /go/bin

# build directories
ADD . /go/src/git.xenonstack.com/util/test-portal
WORKDIR /go/src/git.xenonstack.com/util/test-portal

#Go dep!
#RUN go get -u github.com/golang/dep/...
#RUN dep ensure -update

RUN go install git.xenonstack.com/util/test-portal

ENTRYPOINT /go/bin/test-portal

EXPOSE 8000
