#!/bin/bash
protoc \
		-I. \
		--go_out=. \
		--go-grpc_out=. \
		--openapiv2_out=. \
		--grpc-gateway_out=. \
		--grpc-gateway_opt generate_unbound_methods=true \
		project/project.proto \
		cluster/cluster.proto \
		vm/vm.proto
