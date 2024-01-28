############################
# STEP 1 build executable binary
############################
FROM golang:1.21-alpine AS builder

WORKDIR /go/src/app
COPY . .

RUN go install

############################
# STEP 2 build a small image
############################
FROM alpine
# Copy our static executable.
WORKDIR /root/
COPY --from=builder /go/bin/runic .
COPY --from=builder /go/src/app/etc/.runic.yaml ./.runic.yaml
CMD ["./runic", "server"]
