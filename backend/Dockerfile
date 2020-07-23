FROM golang AS build-stage

WORKDIR /build

ADD go.mod .
ADD go.sum .

RUN go mod download

ADD . .
RUN go build -o moniteur github.com/aueb-cslabs/moniteur/backend
RUN go build -buildmode=plugin -o aueb-plugin.so github.com/aueb-cslabs/moniteur/backend/plugin/aueb

WORKDIR /dist

RUN cp /build/moniteur .
RUN cp /build/aueb-plugin.so .
RUN cp /build/config/* .

FROM alpine

RUN mkdir /production
COPY --from=build-stage /dist/moniteur /production
COPY --from=build-stage /dist/aueb-plugin.so /production
COPY --from=build-stage /dist/calendar.yml /production
COPY --from=build-stage /dist/mapping.yml /production

RUN apk add libc6-compat

WORKDIR /production

CMD ./moniteur --config config/config.yml --cert certificates/moniteur.crt --key certificates/moniteur.key --plugin aueb-plugin.so
