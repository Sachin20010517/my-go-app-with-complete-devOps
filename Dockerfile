FROM golang:1.22 as base

WORKDIR /app

#Having COPY go.mod . & RUN go mod download will be usefull when developer team add new dependencies to the the project. Then it will ne included to the project
COPY go.mod .
RUN go mod download

#Copy the source code on to the docker image
COPY . .

RUN go build -o main .

# Final stage - Distroless image

FROM gcr.io/distroless/base

#copy the main binary which is in /app directory onto the default directory
COPY --from=base /app/main .

EXPOSE 3000

CMD [ "./main" ]











# FROM golang:1.22 as build

# WORKDIR /go/src/app
# COPY . .

# RUN go mod tidy

# RUN go get

# RUN CGO_ENABLED=0 GOOS=linux go build -o /rest-api-app
# # This command will be critical for ensuring that the Go application can run properly in a minimal environment like a Distroless image

# RUN  chmod +x /rest-api-app

# # Final stage (Distroless)
# FROM gcr.io/distroless/static-debian11

# COPY --from=build /rest-api-app /

# EXPOSE 3000

# CMD ["/rest-api-app"]