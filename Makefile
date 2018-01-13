.PHONY: docker clean

twinger:
	go get github.com/ChimeraCoder/anaconda
	go get github.com/microcosm-cc/bluemonday

	go build -o twinger main.go

docker: twinger
	docker build --tag twinger:latest .

clean:
	rm twinger
