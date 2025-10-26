from golang as builder
workdir app
copy go.mod go.sum ./
run go mod download
copy . .
run   CGO_ENABLED=0 go build -o main ./main.go
from alpine
workdir app
copy --from=builder /app/main main
cmd ["./main"]