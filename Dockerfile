FROM openeuler/openeuler:23.03 as BUILDER
RUN dnf update -y && \
    dnf install -y golang && \
    go env -w GOPROXY=https://goproxy.cn,direct

MAINTAINER zengchen1024<chenzeng765@gmail.com>

# build binary
WORKDIR /go/src/github.com/opensourceways/software-package-github-server
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 go build -a -o software-package-github-server .

# copy binary config and utils
FROM openeuler/openeuler:22.03
RUN dnf -y update && \
    dnf in -y shadow && \
    dnf install -y git && \
    groupadd -g 1000 software-package-github-server && \
    useradd -u 1000 -g software-package-github-server -s /bin/bash -m software-package-github-server

USER software-package-github-server

COPY --chown=software-package-github-server --from=BUILDER /go/src/github.com/opensourceways/software-package-github-server/software-package-github-server /opt/app/software-package-github-server
COPY --chown=software-package-github-server softwarepkg/infrastructure/codeimpl/code.sh /opt/app/code.sh
RUN chmod +x /opt/app/code.sh

ENTRYPOINT ["/opt/app/software-package-github-server"]
