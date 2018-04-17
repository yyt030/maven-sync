# 爬虫

### 说明
本脚本是爬取maven仓库中所有的jar包，保存在本地pkg目录中。

### 运行
```
$ go run main.go -h
Usage of main:
  -process int
    	Number of multiple requests to download packages (default 10)
  -redisUrl string
    	Redis url for save downloaded filename (default "localhost:6379")
  -startUrl string
    	Start download url (default "http://repo.maven.apache.org/maven2/")
```
> 爬取过的jar包名字存储在redis中，避免重复下载
