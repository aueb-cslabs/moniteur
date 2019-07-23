.build-moniteur:
	go build -o bin/moniteur github.com/aueb-cslabs/moniteur

.build-aueb-plugin:
	go build -buildmode=plugin -o bin/aueb-plugin.so github.com/aueb-cslabs/moniteur/aueb

.build: .build-moniteur .build-aueb-plugin

default: .build

start: .build
	bin/moniteur