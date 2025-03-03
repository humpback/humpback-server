```shell

docker run -d \
--name=humpback \
--net=host \
--restart=always \
-v /etc/localtime:/etc/localtime \
-e LOCATION=dev \
-e SITE_PORT=8300 \
-e BACKEND_PORT=8301 \
registry.cn-hangzhou.aliyuncs.com/skyler_public/humpback:0.0.1

```