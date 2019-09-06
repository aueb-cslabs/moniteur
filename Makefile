.build-moniteur:
	go build -o bin/moniteur github.com/aueb-cslabs/moniteur

.build-aueb-plugin:
	go build -buildmode=plugin -o bin/aueb-plugin.so github.com/aueb-cslabs/moniteur/aueb

.build-go: .build-moniteur .build-aueb-plugin

.build-vue:
	cd app; npm install; npm run build

.build: .build-go .build-vue

default: .build

test: .build-go
	go test github.com/aueb-cslabs/moniteur
	go test github.com/aueb-cslabs/moniteur/aueb

start: .build-go
	bin/moniteur