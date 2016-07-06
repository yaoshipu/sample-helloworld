#### 准备工作

- 本地安装 Docker ：参考 <a href="https://docs.docker.com/" target="_blank">Docker 官方文档</a>，选择适合操作系统的版本。
- 获得示例的源代码 ：`git clone git@github.com:kirk-enterprise/sample-helloworld.git`
- 示例代码提供 go 和 nodejs 版本，前者的编译需要本地安装 go，参考<a href="https://golang.org/dl" target="_blank"> golang 环境下载</a>。

#### go 版本

helloworld.go 是一个简单的 go WEB 服务程序。

```go
package main

import (
  "fmt"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Handle", r.Method, r.URL)
  fmt.Fprintf(w, "Hello world from my Go program!\n")
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8000", nil)
}
```

容器化的工作包括：

- 将代码交叉编译成适合 linux 运行的二进制。
- 使用 `Dockerfile` 和 `docker build` 命令构建镜像，更多参考 <a href="https://docs.docker.com/engine/getstarted/step_four/" target="_blank">Docker 构建镜像指南</a>。

```bash
[ go ]$ GOOS=linux GOARCH=amd64 go build helloworld.go
[ go ]$ cat Dockerfile
FROM ubuntu
ADD helloworld /
CMD ["/helloworld"]
[ go ]$ docker build -t kirk-test/helloworld-go .
```

使用 `docker run` 运行容器后程序会监听 8000 端口，以 curl 或浏览器可以访问。

```bash
[ go ]$ docker run -d -p 8000:8000 kirk-test/helloworld-go
7a5a78195f2b577cfc5beb086fe4f67694944289b8fecc0f1c7766a2821b2c13
[ go ]$ curl localhost:8000
Hello world from my Go program!

# 执行 `docker rm -f 7a5a78195f2b` 可以清理运行的容器
```

#### nodejs 版本

helloworld.js 是一个简单的 nodejs WEB 服务程序。

```javascript
var http = require('http');

var server = http.createServer(function (request, response) {
  console.log("Handle " + request.method + " " + request.url);
  response.writeHead(200, {"Content-Type": "text/plain"});
  response.end("Hello world from my Node program!\n");
});

server.listen(8000);

console.log("Server running at http://127.0.0.1:8000/");
```

容器化的工作包括：

- 使用 `Dockerfile` 和 `docker build` 命令构建镜像，更多参考 <a href="https://docs.docker.com/engine/getstarted/step_four/" target="_blank">Docker 构建镜像指南</a>。

```bash
[ nodejs ]$ cat Dockerfile
FROM node
ADD helloworld.js /
CMD ["node", "helloworld.js"]
[ nodejs ]$ docker build -t kirk-test/helloworld-node .
```

使用 `docker run` 运行容器后程序会监听 8000 端口，以 curl 或浏览器可以访问。

```bash
[ nodejs ]$ docker run -d -p 8000:8000 kirk-test/helloworld-node
9f77df5feb5ff76ab02792962a8bae4f118fe90a5febf7226ae403a3283a823f
[ nodejs ]$ curl localhost:8000
Hello world from my Node program!
# 执行 `docker rm -f 9f77df5feb5f` 可以清理运行的容器
```
