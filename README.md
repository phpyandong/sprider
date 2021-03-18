＃sprider
这是个golang实现的爬虫项目
＃监听redis
sudo tcpdump ip -n  host 140.143.139.224
```
go ge -u
```
# 如果使用mac编译好文件 itemsave ，在docekr中执行会报|
# standard_init_linux.go:211: exec user process caused "exec format error"
#1 普通并发版
    cd src/sprider
    go run main.go

#2 微服务版
- cd src/sprider/craw/rpcsupport/server
    手动执行
    ```
    go run main.go --port=1234
    go run main.go --port=1235

    ```

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
