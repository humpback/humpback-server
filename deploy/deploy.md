```shell

docker run -d \
--name=humpback \
--net=host \
--restart=always \
-v /etc/localtime:/etc/localtime \
-v /workspace/data/humpback.db:/var/lib/humpback/humpback.db \
-e LOCATION=dev \
-e SITE_PORT=8300 \
-e BACKEND_PORT=8301 \
docker.io/humpbacks/humpback-server:develop

```