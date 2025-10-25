from golang as builder
workdir app
copy go.mod go.sum ./
run go mod download
copy . .
run  go CGO_ENABLED=0 build -o main ./main.go
from alpine
workdir app
copy --from=builder main main
cmd ["./main"]