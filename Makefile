.build-moniteur:
	cd backend && go build -o bin/moniteur github.com/aueb-cslabs/moniteur/backend

.build-aueb-plugin:
	cd backend && go build -buildmode=plugin -o bin/aueb-plugin.so github.com/aueb-cslabs/moniteur/backend/plugin/aueb

.build-go: .build-moniteur .build-aueb-plugin

.build-web:
	cd app && npm install && npm run build

.build: .build-go .build-web

default: .build

test:
	cd backend && go get && go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

start: .build-go
	cd backend && bin/moniteur
