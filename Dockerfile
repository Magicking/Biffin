FROM golang
MAINTAINER Sylvain Laurent

ENV GOBIN $GOPATH/bin

RUN go get -v github.com/ethereum/go-ethereum

ADD vendor /go/src/biffin/vendor
ADD cmd /go/src/biffin/cmd
ADD models /go/src/biffin/models
ADD restapi /go/src/biffin/restapi
ADD internal /go/src/biffin/internal
ADD run.sh /go/src/biffin
WORKDIR /go/src/biffin
RUN go install /go/src/biffin/cmd/biffin-server/main.go
ENTRYPOINT /go/src/biffin/run.sh

# serving HTTP of 8090
EXPOSE 8090
