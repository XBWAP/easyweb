# easyweb
A very simple web server  supporting HTTP and HTTPS 

步骤 1: 更新软件包列表
sudo apt update

步骤 2: 安装 Go
sudo apt install golang-go
这将安装 Go 语言环境，包括 go 命令和 Go 标准库。

步骤 3: 验证 Go 安装
go version
这应该显示刚刚安装的 Go 版本。

步骤 4: 设置 Go 环境
默认情况下，Go 将安装到 /usr/lib/go，但您可能想要设置自己的 Go 环境。创建一个新的目录用于 Go 项目，例如：
mkdir ~/go
然后，将 GOPATH 环境变量设置为指向该目录：
export GOPATH=~/go
您可以将该行添加到 shell 配置文件 (~/.bashrc 或 ~/.zshrc) 中，以使更改永久生效。

步骤 5: 安装附加依赖项（可选）
根据项目的要求，您可能需要安装附加依赖项，例如：
gcc 用于构建使用 C 代码的 Go 程序：
sudo apt install gcc
libssl-dev 用于构建使用 SSL/TLS 的 Go 程序：
sudo apt install libssl-dev
libpq-dev 用于构建使用 PostgreSQL 的 Go 程序：
sudo apt install libpq-dev

简易http服务器

如果 index.html 需要调用该目录下的静态文件，例如 JavaScript 文件、CSS 文件、图片等，可以使用 http.FileServer 函数来实现。

http.FileServer 函数可以将一个目录下的所有文件返回给客户端，例如：


http.Handle("/", http.FileServer(http.Dir("./static")))
这将返回 ./static 目录下的所有文件。

然后，在 index.html 文件中，可以使用相对路径来引用这些静态文件，例如：


<html>
  <head>
    <title>My Web Server</title>
    <link rel="stylesheet" type="text/css" href="styles.css">
  </head>
  <body>
    <h1>Welcome to my web server!</h1>
    <script src="script.js"></script>
    <img src="image.png" alt="image">
  </body>
</html>
在上面的示例中，styles.css、script.js 和 image.png 都是位于 ./static 目录下的静态文件。

当客户端请求 index.html 文件时，http.FileServer 函数将返回该文件的内容，并且当客户端请求这些静态文件时，也将返回这些文件的内容。

完整的示例代码如下：


package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

在上面的代码中，我们使用 http.FileServer 函数将 ./static 目录下的所有文件返回给客户端，并使用 http.HandleFunc 函数将 index.html 文件返回给客户端。

这样，当客户端请求 index.html 文件时，将返回该文件的内容，并且可以正确地加载该文件所需的静态文件。

运行
go run mian.go


使用go build命令生成可执行文件，并指定输出文件名：

go build -o easyweb main.go
这将编译main.go文件并生成一个名为easyweb的可执行文件，可以使用以下命令运行它：

./easyweb



GO简易文件服务器：


如果您想提供./static目录中的所有文件，包括index.html，可以删除自定义处理器函数，并依赖默认文件提供行为：


package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}

这样，./static目录中的所有文件将被HTTP服务器直接提供。


goweb服务器默认加载index.html，并允许加载静态资源


要实现网页默认显示 index.html 文件，并允许 index.html 调用本目录下的 js 和 css 文件，可以使用以下方法：

使用 http.FileServer 处理静态文件

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
这将使得 HTTP 服务器从 ./static 目录中提供所有文件，包括 index.html、js 和 css 文件。

使用 http.HandleFunc 指定默认索引文件
Copy
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
这将使得当用户访问根目录时，自动显示 index.html 文件。


实现网页默认显示index.html文件，并允许index.html调用本目录下的js,css文件，并允许访问该目录下的其他子网页


下面是一个 Go 网页服务器的示例代码，它可以：

当访问根 URL 时，默认显示 index.html 文件
允许 index.html 文件访问同目录下的 js 和 css 文件
允许访问同目录下的其他 HTML 文件和子目录
以下是 main.go 代码：

package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "static/index.html")
		} else {
			http.FileServer(http.Dir("./static")).ServeHTTP(w, r)
		}
	})
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
以下是代码的工作原理：

http.Handle("/", http.FileServer(http.Dir("./static"))) 行设置了一个文件服务器，以便从 ./static 目录中提供文件。
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {...}) 行设置了一个处理函数，以便处理根 URL ("/") 的请求。
在处理函数中，如果请求 URL 正好是 "/"，则它将从 ./static目录中提供index.html文件，使用http.ServeFile`。
如果请求 URL 不是 "/"，则它将使用文件服务器提供请求的文件，从 ./static` 目录和其子目录中。
使用这种设置，当您访问 http://localhost:8080/ 时，它将提供 index.html 文件从 ./static 目录中。index.html 文件可以使用相对 URL 访问同目录下的 js 和 css 文件，例如 <script src="script.js"></script> 和 <link rel="stylesheet" href="style.css">。

您也可以访问 ./static 目录和其子目录中的其他 HTML 文件，例如 http://localhost:8080/about.html 或 http://localhost:8080/subdir/page.html。


正确加载注册系统

你正在使用 Go 语言创建一个 Web 服务器，并想将其编译成一个可执行文件，以便注册为服务，并正确加载 webdata 目录。

要实现这一点，你需要按照以下步骤操作：

将 Go 程序编译成可执行文件：
go build -o webserver main.go
这将在当前目录下创建一个名为 webserver 的可执行文件。

创建服务配置文件（例如 webserver.service），内容如下：
[Unit]
Description=Web 服务器
After=network.target

[Service]
User=<你的用户名>
ExecStart=/path/to/webserver
Restart=always

[Install]
WantedBy=multi-user.target
将 <你的用户名> 替换为实际的用户名，/path/to/webserver 替换为可执行文件的实际路径。

将 webdata 目录复制到与可执行文件相同的位置，以便 Web 服务器可以找到它。

注册服务文件：

sudo systemctl daemon-reload
sudo systemctl enable webserver
sudo systemctl start webserver
这将注册服务并启动它。你可以使用 sudo systemctl status webserver 查看服务状态。

确保 Web 服务器可以正确加载 webdata 目录，使用以下方法：
Instead of using http.FileServer(http.Dir("./webdata")), 使用 http.FileServer(http.Dir("/path/to/webdata")),其中 /path/to/webdata 是 webdata 目录的绝对路径。

或者，你可以使用 os 包来获取当前工作目录，并动态构造 webdata 目录的路径：

package main

import (
	"fmt"
	"net/http"
	"path/filepath"
	"runtime"
)

func main() {
	_, file, _, _ := runtime.Caller(0)
	webDataDir := filepath.Dir(file) + "/webdata"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, webDataDir+"/index.html")
		} else {
			http.FileServer(http.Dir(webDataDir)).ServeHTTP(w, r)
		}
	})
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}这样，Web 服务器将正确加载 webdata 目录，不管当前工作目录是什么。

Let's break down each line of code:

1. package main
定义了这个 Go 语言程序的包名为 main。在 Go 语言中，main 包是程序的入口点。

2. import ( ... )
导入了四个包：

fmt：提供格式化输出函数，例如 Println。
net/http：提供 HTTP 服务器和客户端实现。
path/filepath：提供文件路径操作函数，例如 Dir。
runtime：提供运行时信息和操作函数，例如 Caller。
3. func main() { ... }
定义了程序的入口点函数 main。这个函数将被执行 quando 程序启动。

4. _, file, _, _ := runtime.Caller(0)
使用 runtime.Caller 函数获取当前 goroutine 的调用栈信息。Caller(0) 返回当前 goroutine 的调用栈信息，其中包括文件名、行号、函数名等。这里，我们只关心文件名，所以使用 _ 忽略其他返回值。

5. webDataDir := filepath.Dir(file) + "/webdata"
使用 filepath.Dir 函数获取文件的目录路径，然后追加 /webdata 目录名，得到 webDataDir 变量的值。这个变量将用于存储 Web 应用程序的数据文件。

6. http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { ... })
使用 http.HandleFunc 函数注册一个处理函数，用于处理根 URL ("/") 的请求。这个处理函数将被执行 quando 收到请求时。

7. if r.URL.Path == "/" { ... } else { ... }
检查请求的 URL 路径是否为根 URL ("/"）。如果是，则执行 http.ServeFile 函数，否则执行 http.FileServer 函数。

8. http.ServeFile(w, r, webDataDir+"/index.html")
使用 http.ServeFile 函数将 index.html 文件发送回客户端。这个文件位于 webDataDir 目录下。

9. http.FileServer(http.Dir(webDataDir)).ServeHTTP(w, r)
使用 http.FileServer 函数将 webDataDir 目录下的文件提供给客户端。http.Dir 函数将目录路径转换为 http.FileSystem 对象。

10. fmt.Println("Server started on port 8080")
使用 fmt.Println 函数输出一条信息，表明服务器已经启动并监听端口 8080。

11. http.ListenAndServe(":8080", nil)
使用 http.ListenAndServe 函数启动 HTTP 服务器，监听端口 8080，并将请求分发给注册的处理函数。nil 表示没有特殊的处理函数。


实现简单web服务器，支持注册系统服务，控制台显示
所有网卡ip+端口 

package main

import (
	"fmt"
	"net"
	"net/http"
	"path/filepath"
	"runtime"
)

func getLocalIPs() []string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil
	}
	ips := make([]string, 0)
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		if ip != nil && ip.IsGlobalUnicast() {
			ips = append(ips, ip.String())
		}
	}
	return ips
}

func main() {
	_, file, _, _ := runtime.Caller(0)
	webDataDir := filepath.Dir(file) + "/webdata"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, webDataDir+"/index.html")
		} else {
			http.FileServer(http.Dir(webDataDir)).ServeHTTP(w, r)
		}
	})


	ips := getLocalIPs()
for _, ip := range ips {
	fmt.Println("Server started on " + ip + ":8080")
}
http.ListenAndServe(ips[0]+":8080", nil)
}






