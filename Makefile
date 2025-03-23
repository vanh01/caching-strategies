test:
	docker run --rm -i --network=cs-net --cap-add=SYS_ADMIN grafana/k6:latest-with-browser run - <$(in) > $(out)
.PHONY: test
