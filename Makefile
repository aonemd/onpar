SRC_FILE       = onpar.go
COMPONENTS_DIR = components
BIN_DIR	       = bin

default: clean comp build

comp:
	mkdir -p $(BIN_DIR)
	@for f in $(shell ls ${COMPONENTS_DIR}/*.go); do go build -o $(BIN_DIR) $${f}; done
	find $(COMPONENTS_DIR) ! -name '*.go' -name '*.*' -exec cp {} $(BIN_DIR) \;
clean:
	rm onpar
	rm -rf $(BIN_DIR)
build:
	go build $(SRC_FILE)
stop:
	pkill onpar || true
run:
	onpar&
install: default stop
	cp onpar /usr/local/bin
