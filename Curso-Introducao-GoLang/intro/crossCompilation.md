# Roda para cada sistema operacional sem fazer instalação ou versão específica

GOOS=darwin GOARCH=amd64 go build gorotines.go
GOOS=windows GOARCH=amd64 go build gorotines.go
GOOS=linux GOARCH=amd64 go build gorotines.go