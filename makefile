


_gengo:
	@echo $(USER)
	@docker run -v `pwd`:/defs namely/protoc-all  -i . --with-gateway 	  -f  ./pkg/contrato.proto -l go 
	@sudo chown -R $(USER):$(USER) gen
gen: _gengo


setup:
	docker build -t grpc-gateway ./grpc-gateway


cli:
	evans -p 50051 -r


	#  --grpc-gateway_out=logtostderr=true:./gen/\
		# docker run -v `pwd`:/defs namely/protoc-all  -i .  -f  ./pkg/contrato.proto -l go -o ./pkg/


