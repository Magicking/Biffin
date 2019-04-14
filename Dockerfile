FROM golang
MAINTAINER Sylvain Laurent

ENV GOBIN $GOPATH/bin

ADD vendor /go/src/github.com/Magicking/Biffin/vendor
ADD cmd /go/src/github.com/Magicking/Biffin/cmd
ADD models /go/src/github.com/Magicking/Biffin/models
ADD restapi /go/src/github.com/Magicking/Biffin/restapi
ADD internal /go/src/github.com/Magicking/Biffin/internal
ADD run.sh /go/src/github.com/Magicking/Biffin
WORKDIR /go/src/github.com/Magicking/Biffin
RUN go install /go/src/github.com/Magicking/Biffin/cmd/biffin-server/main.go
ENTRYPOINT /go/src/github.com/Magicking/Biffin/run.sh

# serving HTTP of 8090
EXPOSE 8090
