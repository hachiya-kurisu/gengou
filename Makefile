all: gengou

again: clean all

gengou: gengou.go cmd/gengou/main.go
	go build -o gengou cmd/gengou/main.go

clean:
	rm -f gengou

test:
	go test -cover

cover:
	go test -coverprofile=cover.out
	go tool cover -html cover.out

push:
	got send
	git push github

fmt:
	gofmt -s -w *.go cmd/*/main.go

README.md: README.gmi
	sisyphus -f markdown <README.gmi >README.md

doc: README.md

release: push
	git push github --tags

