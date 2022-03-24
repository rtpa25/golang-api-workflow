#this is a multi stage dockerfile which aims to reduce the weight of this dockerfile which is too high as it contains our golang code and all the dependencies of it even after building the golang binaries 
FROM golang:1.17-alpine AS builder

WORKDIR /app
#first dot ensures all files in the simple-bank dir is copyied and then the second dot ensures that all the stuff is copied into the app dir inside the container
COPY . .

RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.12.2/migrate.linux-amd64.tar.gz | tar xvz

#run phase
FROM alpine
WORKDIR /app
#copies the executables from the builder image specifically from the app main directory and puts in the current working directory which is the . of this new image
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate

COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

#just to communicate with other fellow devs about the port running insidet the container
EXPOSE 8080

CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]