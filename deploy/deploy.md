```shell

docker run -d \
--name=humpback \
--net=host \
--restart=always \
-v /etc/localtime:/etc/localtime \
-e LOCATION=dev \
-e PORT=8550 \
registry.cn-hangzhou.aliyuncs.com/skyler_public/humpback:0.0.1

```