.PHONY: clean

PROGNAME := go-deploy-sample
SRCS := $(wildcard *.go)
OUTDIR := ./build
TARGET := $(OUTDIR)/$(PROGNAME)
GO_MOD_FILE := ./go.mod

build: $(TARGET)
	@echo $(TARGET)

run: $(TARGET)
	@$(TARGET)

clean:
	rm -rf $(OUTDIR)/*

$(TARGET): $(SRCS) $(GO_MOD_FILE)
	@go get
	@go build -o $(TARGET)
