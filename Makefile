PID_FILE = ./build/sampleApi.pid

all: watch

watch: fswatch-install start
	@fswatch -0 handlers controller models db data | xargs -0 -I {} make restart

fswatch-install:
	@if ! type "fswatch" > /dev/null; then brew install fswatch; fi

restart: stop start
	@echo "Restarting server"

stop:
	@kill `cat $(PID_FILE)`
	@rm $(PID_FILE)

start:
	@go build -o build/main ./cmd/
	@touch ${PID_FILE}
	@./build/main & echo $$! > $(PID_FILE)