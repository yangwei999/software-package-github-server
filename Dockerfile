FROM golang:1.18.8 as BUILDER

MAINTAINER zengchen1024<chenzeng765@gmail.com>

# build binary
WORKDIR /go/src/github.com/opensourceways/software-package-github-server
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 go build -a -o software-package-github-server .

WORKDIR /go/src/github.com/opensourceways/software-package-github-server/download
RUN GO111MODULE=on CGO_ENABLED=0 go build -a -o download .

# copy binary config and utils
FROM alpine:3.14
COPY --from=BUILDER /go/src/github.com/opensourceways/software-package-github-server/software-package-github-server /opt/app/software-package-github-server
COPY --from=BUILDER /go/src/github.com/opensourceways/software-package-github-server/download/download /opt/app/download
COPY softwarepkg/infrastructure/codeimpl/code.sh /opt/app/code.sh
RUN chmod +x /opt/app/code.sh && apk update && apk add --no-cache git libc6-compat rpm curl

ENTRYPOINT ["/opt/app/software-package-github-server"]
