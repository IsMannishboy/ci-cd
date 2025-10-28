from golang as builder
workdir /app
copy go.mod go.sum ./
run go mod download
copy . .
cmd ["go","run","main.go"]
