```shell

docker run -d \
--name=humpback \
--net=host \
--restart=always \
-v /etc/localtime:/etc/localtime \
-e LOCATION=dev \
-e SITE_PORT=8300 \
-e BACKEND_PORT=8301 \
a.newegg.org/newegg-docker/humpbacks/humpback-server:develop

```