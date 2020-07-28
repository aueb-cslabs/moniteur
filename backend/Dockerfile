FROM golang AS build-stage

WORKDIR /build

ADD go.* ./

RUN go mod download

ADD . .
RUN go build -o moniteur github.com/aueb-cslabs/moniteur/backend

WORKDIR /dist

RUN cp /build/moniteur .

FROM alpine

RUN mkdir /production
COPY --from=build-stage /dist/moniteur /production

RUN apk add libc6-compat

WORKDIR /production

CMD ./moniteur --config config/config.yml \
    --calendar config/calendar.yml \
    --cert certificates/moniteur.crt \
    --key certificates/moniteur.key
