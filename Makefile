.build-moniteur:
	cd backend && go build -o bin/moniteur github.com/aueb-cslabs/moniteur/backend

.build-aueb-plugin:
	cd backend && go build -buildmode=plugin -o bin/aueb-plugin.so github.com/aueb-cslabs/moniteur/backend/plugin/aueb

.build-go: .build-moniteur .build-aueb-plugin

.build-moniteur-win:
	cd app && npm install && npm run build:win

.build-moniteur-linux:
	cd app && npm install && npm run build:linux

.build-admin-win:
	cd adminPanel && npm install && npm run build:win

.build-admin-linux:
	cd adminPanel && npm install && npm run build:linux

.publish-moniteur-win:
	cd app && npm install && npm run publish:win

.publish-moniteur-linux:
	cd app && npm install && npm run publish:linux

.publish-admin-win:
	cd adminPanel && npm install && npm run publish:win

.publish-admin-linux:
	cd adminPanel && npm install && npm run publish:linux

.build: .build-go .build-moniteur-linux .build-admin-linux

default: .build

publish-win: .publish-moniteur-win .publish-admin-win

publish-linux: .publish-moniteur-linux .publish-admin-linux

test: .build-go
	go test github.com/aueb-cslabs/moniteur/backend
	go test github.com/aueb-cslabs/moniteur/backend/plugin/aueb

start: .build-go
	cd backend && bin/moniteur