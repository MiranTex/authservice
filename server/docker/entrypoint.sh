#!/bin/bash

# if [ ! -f "./go.mod" ]; then
#     go mod init github.com/mirantex/gocontainer 

#     go mod tidy
# fi


CompileDaemon --directory=. --build="go build main.go" --command="./main" 