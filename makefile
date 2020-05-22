


_gengo:
	@echo $(USER)
	@docker run -v `pwd`:/defs namely/protoc-all  -i . --with-gateway 	  -f  ./pkg/contrato.proto -l go 
	@sudo chown -R $(USER):$(USER) gen
gen: _gengo


setup:
	docker build -t generic_api .

run:
	docker run -d -p 8080:8080 -p 50051:50051 generic_api

run-dev:
	docker run -p 8080:8080 -p 50051:50051 generic_api


cli:
	evans -p 50051 -r


	#  --grpc-gateway_out=logtostderr=true:./gen/\
		# docker run -v `pwd`:/defs namely/protoc-all  -i .  -f  ./pkg/contrato.proto -l go -o ./pkg/


