#
# Ubuntu 14.04 with goIrcLog DockerImage
# https://github.com/eternnoir/goIrcLog
#
# Pull base image.
FROM google/golang
MAINTAINER Frank Wang "eternnoir@gmail.com"

RUN go 

ADD . /gopath/src/goIrcLog
