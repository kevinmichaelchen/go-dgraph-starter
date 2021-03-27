.PHONY: cq
cq:
	@$(MAKE) code-quality

.PHONY: code-quality
code-quality:
	@$(MAKE) vet
	@$(MAKE) tidy
	@$(MAKE) imports
	@$(MAKE) fmt
	@$(MAKE) lint

.PHONY: tidy
tidy:
	@echo "${GREEN}✓ Pruning dependencies${NC}\n"
	@env GOPRIVATE=${GOPRIVATE} GO111MODULE=${GO111MODULE} go mod tidy

.PHONY: vet
vet:
	@echo "${GREEN}✓ Checking code for correctness${NC}\n"
	@env GOPRIVATE=${GOPRIVATE} GO111MODULE=${GO111MODULE} go vet ./...

.PHONY: imports
imports:
	@echo "${GREEN}✓ Cleaning up imports${NC}\n"
	@echo "${BLUE}✓ This may take a few seconds...${NC}\n"
	docker run --rm --volume "$(shell pwd):/data" cytopia/goimports -w .

.PHONY: importsv
importsv:
	@echo "${GREEN}✓ Cleaning up imports${NC}\n"
	@echo "${BLUE}✓ This may take a few seconds...${NC}\n"
	docker run --rm --volume "$(shell pwd):/data" cytopia/goimports -v -w .

.PHONY: fmt
fmt:
	@echo "${GREEN}✓ Formatting code${NC}\n"
	docker run --rm --volume "$(shell pwd):/data" cytopia/gofmt -s -w .

.PHONY: lint
lint:
	@echo "${GREEN}✓ Checking code style${NC}\n"
	docker run --rm --volume "$(shell pwd):/data" cytopia/golint ./...

.PHONY: hadolint
hadolint:
	docker run --rm -i hadolint/hadolint < Dockerfile
	docker run --rm -i hadolint/hadolint < Dockerfile-binary