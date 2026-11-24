FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
COPY packages/api/v1/ ./packages/api/v1/
RUN CGO_ENABLED=0 go build -mod=readonly -trimpath -o /out/incident-api ./packages/api/v1/cmd/server

FROM scratch
COPY --from=build /out/incident-api /incident-api
COPY .aws/sqs-backlog.json /catalog/.aws/sqs-backlog.json
ENV INCIDENT_ROOT=/catalog
USER 65532:65532
EXPOSE 8080
ENTRYPOINT ["/incident-api"]
