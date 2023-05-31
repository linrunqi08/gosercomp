#/bin/sh

#
go get github.com/tinylib/msgp
go install github.com/tinylib/msgp
$GOPATH/bin/msgp -o msgp_gen.go -io=false -tests=false -file log.go
