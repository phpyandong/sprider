# sprider
# 如果使用mac编译好文件 itemsave ，在docekr中执行会报|
# standard_init_linux.go:211: exec user process caused "exec format error"

#
- cd src/sprider/craw/rpcsupport/server

    ```
    make build //编译go执行文件
    make docker //生成docker 镜像
    ```
- cd sprider
    ```
    docker-compose up -d
    ```
#启动es
docker run -d --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:5.6
