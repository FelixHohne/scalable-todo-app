#bin/bash

docker build -t todo-app .
docker run -dp 127.0.0.1:8080:8080 todo-app
 curl http://localhost:8080/note/0/

