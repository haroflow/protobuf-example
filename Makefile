.PHONY: protos
protos:
	protoc -I=protos/ --go_out=. protos/addressbook.proto