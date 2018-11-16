# ACME Mock server

This is a mock service for ACME Certificate Server

## Build and Install the Binaries

#### Prerequisite Tools

* [Git](https://git-scm.com/)
* [Go (at least Go 1.11)](http://golang.org)

#### Fetch from Gtihub

```
cd $GOPATH/src/github.com/ffrizzo/
git clone git@github.com:ffrizzo/acme.git

```

#### Building

```
cd acme
dep ensure
go build -o bin/acme cmd/main.go
```

After build run `./bin/acme` from current directory 


### Docker

Building and package on docker image

```
docker build -t ffrizzo/acme .
```

Run new docker container
```
docker run -it --rm --name acme -p 7070:7070 ffrizzo/acme
```