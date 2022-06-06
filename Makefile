BUILD_DIR := bin
CGO_CFLAGS := -Wno-nullability-completeness -Wno-expansion-to-defined
export CGO_CFLAGS

.PHONY: frontend server

all: server frontend

server: |$(BUILD_DIR)
	@echo "Building the server..."
	go build -o bin/servepls ./cmd

frontend:
	@echo "Bundling the frontend..."
	@cd frontend/react-pos && yarn install && yarn build && cd ../..

$(BUILD_DIR):
	@echo "Folder $(BUILD_DIR) does not exist, creating it..."
	mkdir -p $@