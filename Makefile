all:
	export CGO_CFLAGS='-Wno-nullability-completeness -Wno-expansion-to-defined'
	go build -o bin/servepls ./cmd
