.PHONY: all clean

APP_NAME := uiflow_checker
SRCS := main.go go.mod

all: dist/linux-amd64/$(APP_NAME) \
     dist/linux-i386/$(APP_NAME) \
     dist/linux-arm/$(APP_NAME) \
     dist/linux-arm64/$(APP_NAME) \
	 dist/windows/$(APP_NAME) \
	 dist/macos/$(APP_NAME)

dist/linux-amd64/$(APP_NAME): $(SRCS)
	mkdir -p $(dir $@)
	cd $(dir $@); GOOS=linux GOARCH=amd64 go build ../..
dist/linux-i386/$(APP_NAME): $(SRCS)
	mkdir -p $(dir $@)
	cd $(dir $@); GOOS=linux GOARCH=386 go build ../..
dist/linux-arm/$(APP_NAME): $(SRCS)
	mkdir -p $(dir $@)
	cd $(dir $@); GOOS=linux GOARCH=arm go build ../..
dist/linux-arm64/$(APP_NAME): $(SRCS)
	mkdir -p $(dir $@)
	cd $(dir $@); GOOS=linux GOARCH=arm64 go build ../..
dist/windows/$(APP_NAME): $(SRCS)
	mkdir -p $(dir $@)
	cd $(dir $@); GOOS=windows GOARCH=386 go build ../..
dist/macos/$(APP_NAME): $(SRCS)
	mkdir -p $(dir $@)
	cd $(dir $@); GOOS=darwin GOARCH=amd64 go build ../..

clean:
	-@$(RM) -r dist
