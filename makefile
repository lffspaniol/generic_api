

_gengo:
	@docker run -v `pwd`:/defs namely/protoc-all  -i .  -f  ./pkg/contrato.proto -l go -o ./pkg/
	@docker run -v `pwd`:/defs namely/gen-grpc-gateway -f ./pkg/contrato.proto -s TemperatureService -o ./grpc-gateway
gen: _gengo


cli:
	evans -p 50051 -r


	#  --grpc-gateway_out=logtostderr=true:./gen/\
		# docker run -v `pwd`:/defs namely/protoc-all  -i .  -f  ./pkg/contrato.proto -l go -o ./pkg/


