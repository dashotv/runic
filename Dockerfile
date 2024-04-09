ARG GITHUB_TOKEN

############################
# STEP 1a build ui
############################
FROM node:20-alpine as ui-builder
ARG GITHUB_TOKEN

ENV GITHUB_TOKEN $GITHUB_TOKEN
WORKDIR /app/ui
COPY ui/package.json ui/package-lock.json ./
RUN echo //npm.pkg.github.com/:_authToken=$GITHUB_TOKEN > ./.npmrc
RUN echo @dashotv:registry=https://npm.pkg.github.com/ >> ./.npmrc
RUN  --mount=type=cache,target=/app/ui/node_modules npm install
COPY ui/ ./
RUN --mount=type=cache,target=/app/ui/node_modules npm run build

############################
# STEP 1b build go binary
############################
FROM golang:alpine AS builder

WORKDIR /go/src/app
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=bind,source=go.sum,target=go.sum \
  --mount=type=bind,source=go.mod,target=go.mod \
  go mod download

COPY . .
COPY --from=ui-builder /app/static ./static

RUN --mount=type=cache,target=/go/pkg/mod \
  go install

############################
# STEP 2 build a small image
############################
FROM alpine
# Copy our static executable.
WORKDIR /root/
COPY --from=builder /go/bin/runic .
COPY .env.vault .
CMD ["./runic", "server"]
