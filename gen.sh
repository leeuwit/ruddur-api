#!/bin/bash
protoc \
		-I. \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=. \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--grpc-gateway_opt generate_unbound_methods=true \
		--openapiv2_out=. \
		--openapiv2_opt logtostderr=true \
		project/project.proto \
		cluster/cluster.proto \
		vm/vm.proto
