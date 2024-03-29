FROM golang:alpine as builder

RUN mkdir $GOPATH/goodliving/rf-env -p && mkdir /opt/app -p
WORKDIR $GOPATH/goodliving/rf-env

COPY . .

RUN go build && cp rf-env /opt/app/rf-env && cp env_info.xml /opt/app/env_info.xml

FROM alpine 

RUN mkdir /opt/app
WORKDIR /opt/app

COPY --from=builder /opt/app/rf-env /opt/app/rf-env
COPY --from=builder /opt/app/env_info.xml /opt/app/env_info.xml

EXPOSE 8000

CMD ["./rf-env"]