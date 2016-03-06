all:
	go env

get-deps:
	# Getting package deps
	go get -v ./...

vet:
	for f in $(find ./ -name '*.go'); do go tool vet -composites=false -unusedresult=false $f; done

test: vet
	go test
