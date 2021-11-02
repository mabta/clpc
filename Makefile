build: build.web build.clpcd
	@cd ./target && \
	cp -rf ../scripts ./ && \
	tar zcvf clpc.tar.gz  clpc

build.web: internal/cmd/web/web.go
	@mkdir -p ./target/clpc && \
	go build -o ./target/clpc/clpc-web internal/cmd/web/web.go && \
	[ -f ./target/clpc/config.json ] || cp -f ./config.json ./target/clpc

build.clpcd: internal/cmd/clpcd/clpcd.go
	@mkdir -p ./target/clpc && \
	go build -o ./target/clpc/clpcd internal/cmd/clpcd/clpcd.go && \
	[ -f ./target/clpc/config.json ] || cp -f ./config.json ./target/clpc

clear:
	@rm -rf ./target
