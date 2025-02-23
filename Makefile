SHELL=/bin/bash
LAST_TAG=$(shell git tag --list 'v*' --sort version:refname | tail -n 1)
NEXT_TAG=$(shell echo "${LAST_TAG}" | sed -e 's/-hotfix//' | awk -F '.' '{print $$1 "." $$2 "." $$3 + 1}')
NEXT_HOTFIX_TAG=$(shell echo "${LAST_TAG}-hotfix")

release: confirm
	git tag ${NEXT_TAG}
	git push origin ${NEXT_TAG}

hotfix: confirm
	git tag ${NEXT_HOTFIX_TAG}
	git push origin ${NEXT_HOTFIX_TAG}

write-version:
	echo $(NEXT_TAG) > cmd/mkgo/version

confirm:
	@if [[ -z "$(CI)" ]]; then \
		read -p "⚠ Are you sure? [y/N] > " -r ; \
		if [[ ! $$REPLY =~ ^[Yy]$$ ]]; then \
			printf "Stopping" ; \
			exit 1; \
		else \
			exit 0; \
		fi \
	fi

build: write-version
	go build -o ${GOPATH}/bin/mkgo ./cmd/mkgo
