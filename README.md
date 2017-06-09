# Go-Simple-Flatbuffer

## How To Run

```bash
# clone into GOPATH
git clone git@gitlab.com:kurio/flatbuffer-go-research.git
# CD into project directory
glide install -v

# RUN
go run main.go

```

## To BenchMark

HTTP Benchmark
```bash
# HTTP Request Benchmark

cd users/delivery/http

go test -bench=.
```

De/Serializing Benchmark
```bash
# De/Serializing Benchmark

cd users/delivery/http/fbs

go test -bench=.
```
