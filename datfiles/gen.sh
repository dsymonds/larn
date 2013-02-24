#! /bin/bash -e

# Requires
#   go install github.com/dsymonds/goembed

goembed -package datfiles -var Help < larn.help > help.go
goembed -package datfiles -var Mazes < larnmaze > maze.go
gofmt -w -l {help,maze}.go
