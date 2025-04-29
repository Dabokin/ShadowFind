#!/bin/bash
cd ../core
GOOS=linux GOARCH=amd64 go build -o ../bin/pantherfuzz-core ./main.go
echo "[+] Go-библиотека скомпилирована в bin/pantherfuzz-core"