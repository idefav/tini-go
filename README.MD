# tini Go 版本实现

C语言版本: https://github.com/krallin/tini

## 1. 构建
```bash
make
```

## 2. 测试镜像构建
```bash
cd ./demo
docker build --no-cache -t idefav/tini-go-demo .
```

## 3. 测试镜像运行
```bash
docker run --rm --name="tini-go-demo2" idefav/tini-go-demo
```

