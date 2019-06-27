FROM golang:alpine as builder

RUN mkdir /opt/app
WORKDIR /opt/app

COPY . .

RUN go build 

FROM alpine 

RUN mkdir /opt/app
WORKDIR /opt/app

COPY --from=builder /opt/app/rf-env /opt/app/rf-env
COPY --from=builder /opt/app/env_info.xml /opt/app/env_info.xml

EXPOSE 8000

CMD ["./rf-env"]