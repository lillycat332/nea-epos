all:
	CGO_CFLAGS='-Wno-nullability-completeness -Wno-expansion-to-defined'
	go build -o build/serve