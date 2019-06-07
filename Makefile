GO=go
MAIN=./cmd/nsrl_checker/*.go
GOOS?=unkos
GOARCH?=unkarch
OUTPUT=./bin/$(GOOS)/$(GOARCH)/nsrl_checker
RM=rm
DEP=dep
VENDOR=./vendor

#dependencies: $(VENDOR)
#	$(DEP) ensure

#build: dependencies
build: 
	$(GO) build -o $(OUTPUT) $(MAIN)

clean:
	$(RM) -rf ./bin 	