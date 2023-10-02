FROM 3dsinteractive/golang:1.20-alpine3.16-librdfkafka1.9.2

WORKDIR /go/src/bitbucket.org/automationworkshop/main
ADD . /go/src/bitbucket.org/automationworkshop/main
RUN go build -mod vendor -tags "musl static_all" .

# ================
FROM 3dsinteractive/alpine:3.16

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=0 /go/src/bitbucket.org/automationworkshop/main/main /main

ADD ./entrypoint.sh /entrypoint.sh

RUN adduser -u 1001 -D -s /bin/sh -G ping 1001
RUN chown 1001:1001 /entrypoint.sh
RUN chown 1001:1001 /main

RUN chmod +x /entrypoint.sh
RUN chmod +x /main

USER 1001

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]
