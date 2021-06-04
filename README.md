# protobuf-example
Learning Protocol Buffers from Google, with Go.

https://developers.google.com/protocol-buffers/docs/gotutorial

## Requirements

- go1.16
- protoc. Download [here](https://developers.google.com/protocol-buffers/docs/downloads). Add the binary to your PATH.
- protoc-gen-go. Install by running this command outside a go module:
  ```
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  ```

## Running the example

With `go`, `protoc` and `protoc-gen-go` installed:
```
git clone https://github.com/haroflow/protobuf-example
cd protobuf-example

# Generate .go from protobuf schema
make

# Run the example
go run .
```