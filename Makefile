release: build.web build.clpcd
	@cd ./target && \
	tar zcvf clpc.tar.gz  clpc

release.full: build.web build.clpcd copy.cfg
	@cd ./target && \
	cp -rf ../scripts ./clpc && \
	tar zcvf clpc.tar.gz  clpc

copy.cfg: ./config.json
	[ -f ./target/clpc/config.json ] || cp -f ./config.json ./target/clpc 

build.web: internal/cmd/web/web.go
	@mkdir -p ./target/clpc && \
	go build -o ./target/clpc/clpc-web internal/cmd/web/web.go 

build.clpcd: internal/cmd/clpcd/clpcd.go
	@mkdir -p ./target/clpc && \
	go build -o ./target/clpc/clpcd internal/cmd/clpcd/clpcd.go 

clear:
	@rm -rf ./target
