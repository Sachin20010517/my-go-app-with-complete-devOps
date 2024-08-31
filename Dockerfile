FROM golang:1.22 as base

WORKDIR /go/src/app

#Having COPY go.mod . & RUN go mod tidy will be usefull when developer team add new dependencies to the the project. Then it will ne included to the project
COPY go.mod go.sum ./
RUN go mod tidy 
#RUN go get

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /rest-api-app
# This command will be critical for ensuring that the Go application can run properly in a minimal environment like a Distroless image

RUN  chmod +x /rest-api-app

# Final stage (Distroless)
FROM gcr.io/distroless/static-debian11

COPY --from=base /rest-api-app /

EXPOSE 3000

CMD ["/rest-api-app"]