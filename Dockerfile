FROM golang:1.14.5

# Don't run tests as root so we can play with permissions
RUN useradd --create-home --user-group app

ENV GOPACKAGE github.com/Barzahlen/gotidus

ADD . /go/src/$GOPACKAGE
RUN chown -R app /go

WORKDIR /go/src/$GOPACKAGE

USER app
RUN go build
