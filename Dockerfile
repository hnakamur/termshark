# syntax = docker/dockerfile:experimental

# use golang base image
ARG GO_VERSION=1.13.3
FROM golang:${GO_VERSION}-buster

# install nfpm
ARG NFPM_VERSION=v1.0.0-beta3
ADD https://github.com/goreleaser/nfpm/releases/download/${NFPM_VERSION}/nfpm_amd64.deb /tmp/
RUN dpkg -i /tmp/nfpm_amd64.deb

# Build executable to ./termshark
ARG TERMSHARK_VERSION
COPY . /src
WORKDIR /src
RUN go build -tags netgo -ldflags "-X github.com/gcla/termshark.Version=${TERMSHARK_VERSION}" ./cmd/termshark/

# build deb and rpm packages
RUN tar cf - termshark | gzip -9 > termshark.linux-amd64.tar.gz
RUN nfpm pkg --target termshark.amd64.deb
RUN nfpm pkg --target termshark.x86_64.rpm

## make release
ARG GITHUB_RELEASE_VERSION=v0.7.2
RUN curl -sSL https://github.com/hnakamur/github-release/releases/download/$GITHUB_RELEASE_VERSION/github-release.linux-amd64.tar.gz | tar zxf - -C /usr/local/bin
ARG GITHUB_USER=hnakamur
ARG GITHUB_REPO=termshark
RUN --mount=type=secret,id=github_token,target=/github_token \
  github-release release \
    -s $(cat /github_token) \
    -t v$TERMSHARK_VERSION \
    -c $(git rev-list -n 1 v$TERMSHARK_VERSION) \
  && github-release upload \
    -s $(cat /github_token) \
    -t v$TERMSHARK_VERSION \
    -f termshark.linux-amd64.tar.gz \
    -n termshark.linux-amd64.tar.gz \
  && github-release upload \
    -s $(cat /github_token) \
    -t v$TERMSHARK_VERSION \
    -f termshark.amd64.deb \
    -n termshark.amd64.deb \
  && github-release upload \
    -s $(cat /github_token) \
    -t v$TERMSHARK_VERSION \
    -f termshark.x86_64.rpm \
    -n termshark.x86_64.rpm
