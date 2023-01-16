FROM docker.io/golang:1.19.5-alpine3.17 AS builder
WORKDIR /builder

ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 LDFLAGS="-w -s"
ENV REPOUSER=kazimsarikaya REPOPROJECT=citrixadccertmanagerwebhook REPOHOST=github.com

COPY . .
RUN apk add make gawk jq curl git
RUN make build


FROM scratch
COPY --from=builder /builder/bin/citrixadccertmanagerwebhook /bin/citrixadccertmanagerwebhook
ENTRYPOINT ["/bin/citrixadccertmanagerwebhook"]
