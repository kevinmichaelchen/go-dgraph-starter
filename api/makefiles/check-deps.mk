.PHONY: check-deps
check-deps:
	env GOPRIVATE=${GOPRIVATE} go list -u -m -json all | docker run -i --rm psampaz/go-mod-outdated -update -direct