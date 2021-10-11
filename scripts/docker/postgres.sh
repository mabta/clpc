docker run \
	--name clpc_postgres \
	--restart=always \
	-e POSTGRES_PASSWORD=9mqexnfukpeeC6Ze5fXgkBE6xFHwz5Zuzf27BvAN \
	-e POSTGRES_USER=clpc \
	-e POSTGRES_DB=clpc \
	-p 127.0.0.1:15432:5432 \
	-v /var/docker/clpc_postgres:/var/lib/postgresql/data \
	-d postgres:alpine
