
all:
	go build
	mkdir -p target/bin
	cp hello-server target/bin/launch