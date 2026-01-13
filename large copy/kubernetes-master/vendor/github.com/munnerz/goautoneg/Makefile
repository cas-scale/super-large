include $(GOROOT)/src/Make.inc

TARG=bitbucket.org/ww/goautoneg
GOFILES=autoneg.go

include $(GOROOT)/src/Make.pkg

format:
	gofmt -w *.go

docs:
	gomake clean
	godoc ${TARG} > README.txt
# ID-1768294487-8a9351b9
