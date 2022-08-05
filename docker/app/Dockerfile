FROM golang:1.19-alpine

WORKDIR /app

COPY . .
RUN go mod download


RUN cd cmd/app && go build -o /application

EXPOSE 8080

CMD [ "/application" ]