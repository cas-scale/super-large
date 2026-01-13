include $(GOROOT)/src/Make.inc

TARG=bitbucket.org/ww/goautoneg
GOFILES=autoneg.go

include $(GOROOT)/src/Make.pkg

format:
	gofmt -w *.go

docs:
	gomake clean
	godoc ${TARG} > README.txt
# ID-1768294467-5c4b28e7
