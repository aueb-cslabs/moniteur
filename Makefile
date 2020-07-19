.build-moniteur:
	cd backend && go build -o bin/moniteur github.com/aueb-cslabs/moniteur/backend

.build-aueb-plugin:
	cd backend && go build -buildmode=plugin -o bin/aueb-plugin.so github.com/aueb-cslabs/moniteur/backend/plugin/aueb

.build-go: .build-moniteur .build-aueb-plugin

.build-web:
	cd app && npm install && npm run build

.build: .build-go .build-web

default: .build

test: .build-go
	go test github.com/aueb-cslabs/moniteur/backend
	go test github.com/aueb-cslabs/moniteur/backend/plugin/aueb

start: .build-go
	cd backend && bin/moniteur
