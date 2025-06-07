# run the dcoker container

docker run --name all-dev-tools-docker-2 -e POSTGRES_USER=div -e POSTGRES_PASSWORD=div123 -p 5432:5432 -d postgres

# run the index.html (can use thunderclient from vscode)

# run go server

cd go-backend
go run main.go

