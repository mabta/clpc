build: build.web build.clpcd
	@tar zcvf ./target/clpc.tar.gz  ./target/clpc

build.web: internal/cmd/web/web.go
	@mkdir -p ./target/clpc && \
	go build -o ./target/clpc/clpc_web internal/cmd/web/web.go && \
	[ -f ./target/clpc/config.json ] || cp -f ./config.json ./target/clpc

build.clpcd: internal/cmd/clpcd/clpcd.go
	@mkdir -p ./target/clpc && \
	go build -o ./target/clpc/clpcd internal/cmd/clpcd/clpcd.go && \
	[ -f ./target/clpc/config.json ] || cp -f ./config.json ./target/clpc

clear:
	@rm -rf ./target