

_gengo:
	@protoc -I. \
	 --go_out=plugins=grpc,paths=source_relative:./gen/ \
	 --grpc-gateway_out=logtostderr=true:. \
	 contratc.proto
gen: _gengo


cli:
	evans -p 50051 -r