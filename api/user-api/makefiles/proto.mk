.PHONY: protoc
protoc:
	docker run --rm -v $(shell pwd):/defs namely/protoc-all -l go -i ./proto -d ./proto/myorg/user/v1 -o ./pkg/pb --go-source-relative

.PHONY: protodoc
protodoc:
	cd proto && \
	protodoc --directory=. --languages="Go" --parse="service,message" --title=MyOrg --output=README.md