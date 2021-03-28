set -ex
rm -rf go.mod go.sum
go mod init github.com/riba2534/torrent-client
go mod tidy
go build -o main