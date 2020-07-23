FROM golang AS build-stage

WORKDIR /build

ADD go.* ./

RUN go mod download

ADD . .
RUN go build -o moniteur github.com/aueb-cslabs/moniteur/backend
RUN go build -buildmode=plugin -o aueb-plugin.so github.com/aueb-cslabs/moniteur/backend/plugin/aueb

WORKDIR /dist

RUN cp /build/moniteur .
RUN cp /build/aueb-plugin.so .

FROM alpine

RUN mkdir /production
COPY --from=build-stage /dist/moniteur /production
COPY --from=build-stage /dist/aueb-plugin.so /production

RUN apk add libc6-compat

WORKDIR /production

CMD ./moniteur --config config/config.yml \
    --calendar config/calendar.yml \
    --cert certificates/moniteur.crt \
    --key certificates/moniteur.key  \
    --plugin aueb-plugin.so
