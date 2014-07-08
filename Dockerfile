#
# Ubuntu 14.04 with goIrcLog DockerImage
# https://github.com/eternnoir/goIrcLog
#
# Pull base image.
FROM google/golang
MAINTAINER Frank Wang "eternnoir@gmail.com"

RUN go get github.com/thoj/go-ircevent
RUN go get code.google.com/p/go-sqlite/go1/sqlite3

ADD . /gopath/src/goIrcLog
WORKDIR /gopath/src/goIrcLog
RUN go build

VOLUME ["/config"]
VOLUME ["/db"]

CMD ["./goIrcLog","/config/config.json"]

