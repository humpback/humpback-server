FROM dev-harbor.newegg.org/base/alpine:3

LABEL maintainer="skyler.w.yang"

RUN  mkdir -p /workspace/config

COPY ./backend/config/*.yaml /workspace/config

COPY ./front/projects/web/dist /workspace/html/web

COPY ./backend/humpback /workspace/

WORKDIR /workspace

CMD ["./humpback"]
