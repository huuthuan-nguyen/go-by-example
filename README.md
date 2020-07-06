# How to run Go with Docker?

```bash
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.14 go run hello.go
```

# How to build Go app with Docker?

```bash
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.14 go build hello.go
```
