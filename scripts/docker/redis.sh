docker run \
	--name clpc_redis \
	--restart=always \
	-p 127.0.0.1:16379:6379 \
	-v /var/docker/clpc_redis:/data \
	-d redis:alpine
