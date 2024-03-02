# gin框架学习

## 一、使用net包搭建web服务器

main.go

```go
package main

import (
	"fmt"
	"net/http"
	"os"
)

// 编写一个方法使之在浏览器输入相应的url便可得到浏览器上响应的页面
func sayHello(w http.ResponseWriter, r *http.Request) {
	//将响应在网页的东西写在一个文件里面，通过读取文件里的内容进行响应
	b, _ := os.ReadFile("./hello.txt") //读hello.txt这个文件
	_, _ = fmt.Fprintln(w, string(b))  //进行输出操作
}
func main() {
	fmt.Println("hello world")
	//如果你向浏览器输入/hello我就给你执行sayHello这个函数
	http.HandleFunc("/hello", sayHello)
	//
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http Server failed,err:%v\n", err)
		return
	}

}

```

hello.txt

```go
<h1 style='color:orange'> hello golang!<h1>
<h1>How are you golang! <h1>
<img id='i1' src='https://tse3-mm.cn.bing.net/th/id/OIP-C.xczgV5hOfPJlurtsgWHoqgAAAA?rs=1&pid=ImgDetMain'>
<button id='b1'>点我</button>
<script>
document.getElementById('b1').onclick = function(){
    document.getElementById("i1").src='https://ts1.cn.mm.bing.net/th/id/R-C.4a1c9198b9a745cce043a25b32b629b7?rik=q1h5vuqxdKsqkA&riu=http%3a%2f%2fn.sinaimg.cn%2fsinakd20123%2f83%2fw803h880%2f20200813%2f9bd2-ixreehp6472485.jpg&ehk=0pNKO48qg7f%2bBbs6LMqiPPsgzVMxavgL0vAcAJIVvzc%3d&risl=&pid=ImgRaw&r=0'
}
</script>

```

## 二、gin的初体验

### 1.什么是gin?

一个go语言的web框架

### 2.如何得到gin?(要保证你的go版本可以使用go mod)

```go
go get -u github.com/gin-gonic/gin
```

### 3.案例初体验

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context) {//返回一个json格式的字符串
	c.JSON(200, gin.H{
		"message": "hello golang",
	})
}
func main() {
	r := gin.Default() //返回默认的路由引擎

	//指定用户使用GET请求访问/hello时，执行sayHello这个函数
	r.GET("/hello", sayHello)
	//不采用restful风格的写法
	//r.GET("/book",...)
	//r.GET("/create_book",...)
	//r.GET("/update_book",...)
	//r.GET("/delete_book",...)

    //restful风格:推荐使用，行业规范这里要打开postman用get,post delete....等方法去请求http://127.0.0.1:9090/book来查看是否可以获取json字符串
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	//启动服务
	r.Run(":9090")
}

```

## 三、模板引擎

1.基于http的模板引擎内容

首先创建一个**tmpl模板文件**其实就是一个html文件

hello.**tmpl**

```html
<!DOCTYPE html>
<html lang="zh-CN">
<head>
  <title>hello</title>
</head>
<body>
  <h1>Hello{{.}}</h1>
</body>
</html>
```

注意：{{.}}表示要传入的对象

然后编写一个main.go开始进行操作

```go
package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// 遇事不决，写注释
func sayHello(w http.ResponseWriter, r *http.Request) {
	//2.解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse template fail err: %v", err)
		return
	}
	//3.渲染模板
	name := "小王子"
	err = t.Execute(w, name)
	if err != nil {
		fmt.Println("execution failed err: %v", err)
		return
	}

}
func main() {
	//通过http的方式
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("http server start fail err: %v", err)
		return
	}
}

```

使用gin框架对于tmpl进行渲染

main.go

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default() //创建一个默认的路由
	//r.LoadHTMLFiles("./templates/posts/index.tmpl","./templates/users/index.tmpl") //模板的解析
	r.LoadHTMLGlob("./templates/**/*") //表示对templates下面所有的文件进行一个加载

	r.GET("/posts/index", func(c *gin.Context) {
		//http请求
		c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{ //模板渲染
			"title": "posts/index.com",
		})
	})
```

templates/posts/index.tmpl

```html
{{define "posts/index.tmpl"}}
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>posts/index</title>
</head>
<body>
    {{.title}}
</body>
</html>
{{end}}

```

## 四、gin框架学习

### 1、gin处理json的数据

main.go

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		//方法1：使用map,key是string类型的，后面的值是任意类型的数据
		//data := map[string]interface{}{
		//	"name": "小王子",
		//	"msg":  "hello world",
		//	"age":  18,
		//}
		data := gin.H{"name": "小王子", "message": "hello world", "age": 18}
		c.JSON(http.StatusOK, data)
	})

	//方法2：使用结构体,使用tag来定制化操作，让自己返回的json以自己想要的数据字段返回出去
	type msg struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Message string `json:"message"`
	}
	r.GET("/ajson", func(c *gin.Context) {
		data := msg{
			Name:    "小明",
			Age:     23,
			Message: "我是超人",
		}
		c.JSON(http.StatusOK, data)
	})

	r.Run(":9090")

}

```

### 2、query参数

main.go

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// querystring
func main() {

	r := gin.Default() //创建一个默认的路由
	
	//GETi请求url后面的是query参数是querysting
	//key=value格式，多个key-value值用&连接
	//eq:/web/qurey=小王子&age=78

	r.GET("/web", func(c *gin.Context) {
		//获取浏览器那边发请求qurey string参数
		//第一种
		name := c.Query("query") //通过query获取请求中携带的query参数
		age := c.Query("age")    //通过query获取请求中携带的query参数

		//第二种,没有传值的话就是默认sombody了
		//name := c.DefaultQuery("query", "somebody")
		//第三种
		//name, ok := c.GetQuery("query")//根据取不到就返回第二个false值
		//if !ok {
		//	//取不到
		//	name = "somebody"
		//}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.Run(":9090")
}
```

### 3、form表单提交

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取form表单提交的参数
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./login.html", "./index.html")
	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	//login post的请求，一次请求对应一个响应
	r.POST("/login", func(c *gin.Context) {
		//获取form表单提交的数据
		//username := c.PostForm("username")
		//password := c.PostForm("password") //取到就返回值，取不到就是返回空
		//username := c.DefaultPostForm("username", "somebody")
		//password := c.DefaultPostForm("password", "***")
		username, ok := c.GetPostForm("username")
		if !ok {
			username = "sb"
		}

		password, _ := c.GetPostForm("password")

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Name":     username,
			"Password": password,
		})
	})
	r.Run(":9090")
}

```

### 4、gin获取uri参数

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 获取请求的URI(path)参数,返回的都是字符类型
//注意url的匹配不要冲突
func main() {
	r := gin.Default()

	r.GET("/:name/:age", func(c *gin.Context) {
		//获取路径参数
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})

	})

	r.GET("/blog/:year/:month", func(c *gin.Context) {
		year := c.Param("year")
		month := c.Param("month")
		c.JSON(http.StatusOK, gin.H{
			"year":  year,
			"month": month,
		})
	})
	
	r.Run(":9090")
}

```

### 5、gin参数绑定

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserInfo struct {
	Username string `form:"username" json:"user"`
	Password string `form:"password" json:"pwd"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/user", func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")
		//u := UserInfo{
		//	username: username,
		//	password: password,
		//}
		//以下代码就是对上面注释的进行简化替代操作
		var u UserInfo //声明一个UserInfo类型的变量u
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("拿到的数据是：%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//使用form表单的方式发送post请求
	r.POST("/form", func(c *gin.Context) {
		var u UserInfo //声明一个UserInfo类型的变量u
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("拿到的数据是：%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}

	})
	//在postman种使用json格式发送post请求
	r.POST("/json", func(c *gin.Context) {
		var u UserInfo //声明一个UserInfo类型的变量u
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("拿到的数据是：%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}

	})
	r.Run(":9090")

}

```

### 6、文件上传

上传单个文件

main.go

```go
package main
//上传单个文件
import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {
	r := gin.Default()
    //加载html文件就是写了一个html用post提交操作
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		//从请求中读取文件
		f, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//将读取到的文件保存在本地
			//dst := fmt.Sprintf("./%s",f.Filename)
			dst := path.Join("./", f.Filename)
			c.SaveUploadedFile(f, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.Run(":8080")

}

```

index.html

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>index</title>
</head>
<body>
<form action="/upload" method="post" enctype="multipart/form-data">
  <input type="file" name="f1">
  <input type="submit" value="上传">
</form>
</body>
</html>
```



上传多个文件

```go
func main() {
	router := gin.Default()
	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["file"]

		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
			// 上传文件到指定的目录
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})
	router.Run()
}

```

### 7.请求重定向操作

main.go

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//http转发
	r.GET("/index", func(c *gin.Context) {
		//	c.JSON(http.StatusOK, gin.H{
		//		"status": "ok",
		//	})
		//进行重定向操作
		c.Redirect(http.StatusMovedPermanently, "http://www.bing.com")
	})
//gin 路由转发
	r.GET("/a", func(c *gin.Context) {
		//跳转到b的路由处理函数
		c.Request.URL.Path = "/b" //把请求的URI修改
		r.HandleContext(c)        //继续后续的处理

	})

	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})

	r.Run(":8080")
}

```

### 8、gin路由和路由组

main.go

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	//访问/index的GET请求会走这一条处理逻辑
	//路由
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})

	})
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	//Any可以处理我们刚看到的所有请求，当请求很多的时候
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			c.JSON(http.StatusOK, gin.H{"method": "GET"})
		case http.MethodPost:
			c.JSON(http.StatusOK, gin.H{"method": "POST"})
			//.....
		}
	})
	//NO ROUET,访问不存在的页面
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "没有找到这个页面"})
	})

	////时评的首页和详情页
	//r.GET("/video/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
	//})
	////商城的首页和详情页
	//r.GET("/shop/index", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "/shop/index"})
	//})

	//路由组的组，便于处理多个业务线
	videoGroup := r.Group("video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/index"})
		})

		videoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/xx"})
		})
		videoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"msg": "/video/oo"})
		})
	}

	r.GET("/shop/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "/shop/index"})
	})

	r.GET("/shop/xx", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "/shop/index"})
	})

	r.GET("/shop/oo", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "/shop/index"})
	})
    //路由组也是支持嵌套的
    shopGroup := r.Group("/shop")
    {
        shopGroup.GET("/index",func(c *gin.Context{...}))
        shopGroup.GET("/cart",func(c *gin.Context{...}))
        shopGroup.POST("/checkout",func(c *gin.Context{...}))
        //嵌套路由组
        xx := shopGroup.Group("xx")
        xx.GET("/OO",func(c *gin.Context){...})
    }
	r.Run(":9090")
}

```

### 9、gin中间件

#### 1)介绍

Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等。

#### 2）定义中间件

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// handlerFuncs
func indexHandler(c *gin.Context) {
	fmt.Println("index")
	name, ok := c.Get("name") //跨中间件取值
	if !ok {
		name = "匿名用户"
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

// 定义一个中间件,统计处理函数的耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in ..")
	//计时
	start := time.Now()
	c.Next() //调用后续的处理函数
	//c.Abort()//阻止后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out ..")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ..")
	c.Set("name", "qimi") //设置值
	//c.Abort() //调用后续的处理函数
	//return
	fmt.Println("m2 out ..")
}

func outMiddleware(docheck bool) gin.HandlerFunc {
	//连接数据库
	//或者一些其他的准备工作
	return func(c *gin.Context) {
		if docheck {
			//是否登录的判断
			//if 是登录用户
			c.Next()
			//else
			//c.Abort()
		} else {
			c.Next()
		}
	}
}
func main() {
	r := gin.Default()
	r.Use(m1, m2, outMiddleware(true)) //全局注册中间件函数

	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})

	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})

	////路由组注册中间件1
	//xxGroup := r.Group("/xx",outMiddleware(true))
	//{
	//	xxGroup.GET("/index", func(c *gin.Context) {
	//		c.JSON(http.StatusOK, gin.H{
	//			"msg":"xxGroup",
	//		})
	//	})
	//
	//}
	//
	//
	////路由组注册中间件2
	//xx2Group := r.Group("/xx2")
	//xx2Group.Use(outMiddleware(true))
	//{
	//	xx2Group.GET("/index", func(c *gin.Context) {
	//		c.JSON(http.StatusOK, gin.H{
	//			"msg":"xx2Group",
	//		})
	//	})
	//
	//}
	r.Run(":9090")
}

```

中间件注意事项

gin默认中间件

`gin.Default()`默认使用了`Logger`和`Recovery`中间件，其中：

- `Logger`中间件将日志写入`gin.DefaultWriter`，即使配置了`GIN_MODE=release`。
- `Recovery`中间件会recover任何`panic`。如果有panic的话，会写入500响应码。

如果不想使用上面两个默认的中间件，可以使用`gin.New()`新建一个没有任何默认中间件的路由。

#### gin中间件中使用goroutine

当在中间件或`handler`中启动新的`goroutine`时，**不能使用**原始的上下文（c *gin.Context），必须使用其只读副本（`c.Copy()`）

# GORM的学习

### 1.安装与使用

安装

```go
go get -u github.com/jinzhu/gorm
或者使用 go get -u gorm.io/gorm
```

### 2、连接数据库

连接不同的数据库都需要导入对应数据的驱动程序，`GORM`已经贴心的为我们包装了一些驱动程序，只需要按如下方式导入需要的数据库驱动即可：

把mysql连接的驱动download下来

```go
go get -u github.com/jinzhu/gorm/dialects/mysql 
```



```go
import _ "github.com/jinzhu/gorm/dialects/mysql"
// import _ "github.com/jinzhu/gorm/dialects/postgres"
// import _ "github.com/jinzhu/gorm/dialects/sqlite"
// import _ "github.com/jinzhu/gorm/dialects/mssql"
```

#### -1.连接MySQL

```go
import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
  db, err := gorm.Open("mysql", "user:password@(localhost)/dbname?charset=utf8mb4&parseTime=True&loc=Local")
  defer db.Close()
}
```

#### -2.连接PostgreSQL

基本代码同上，注意引入对应`postgres`驱动并正确指定`gorm.Open()`参数。

```go
import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
  db, err := gorm.Open("postgres", "host=myhost port=myport user=gorm dbname=gorm password=mypassword")
  defer db.Close()
}
```

#### -3.连接Sqlite3

基本代码同上，注意引入对应`sqlite`驱动并正确指定`gorm.Open()`参数。

```go
import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
  db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
  defer db.Close()
}
```

#### -4.连接SQL Server

基本代码同上，注意引入对应`mssql`驱动并正确指定`gorm.Open()`参数。

```go
import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mssql"
)

func main() {
  db, err := gorm.Open("mssql", "sqlserver://username:password@localhost:1433?database=dbname")
  defer db.Close()
}
```

### 3.实例操作

```go
package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//使用gorm进行简单书数据库连接并进行增删查改的工作

// UeriNFO
type UserInfo struct {
	ID     int
	Name   string
	Gender string
	Hobby  string
}

func main() {
	//连接mysql数据库
	db, err := gorm.Open("mysql", "root:123456@(192.168.20.10:3306)/db1?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功!")
	defer db.Close()
	//创建表,自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	//创建数据行
	//u1 := UserInfo{1, "boat", "男", "篮球"}
	//db.Create(u1)
	//查询
	var u UserInfo
	db.First(&u) //查询表中第一条数据保存到u中
	fmt.Printf("u:%#v\n", u)
	//更新
	db.Model(&u).Update("hobby", "双色球ss")
	fmt.Printf("u:%#v\n", u)
	//删除
	db.Delete(&u)
}

```

4.模型定义

## GORM Model定义

在使用ORM工具时，通常我们需要在代码中定义模型（Models）与数据库中的数据表进行映射，在GORM中模型（Models）通常是正常定义的结构体、基本的go类型或它们的指针。 同时也支持`sql.Scanner`及`driver.Valuer`接口（interfaces）。

### gorm.Model

为了方便模型定义，GORM内置了一个`gorm.Model`结构体。`gorm.Model`是一个包含了`ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`四个字段的Golang结构体。

```go
// gorm.Model 定义
type Model struct {
  ID        uint `gorm:"primary_key"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt *time.Time
}
```

你可以将它嵌入到你自己的模型中：

```go
// 将 `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`字段注入到`User`模型中
type User struct {
  gorm.Model
  Name string
}
```

当然你也可以完全自己定义模型：

```go
// 不使用gorm.Model，自行定义模型
type User struct {
  ID   int
  Name string
}
```

### 模型定义示例

```go
type User struct {
  gorm.Model
  Name         string
  Age          sql.NullInt64
  Birthday     *time.Time
  Email        string  `gorm:"type:varchar(100);unique_index"`
  Role         string  `gorm:"size:255"` // 设置字段大小为255
  MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
  Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
  Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
  IgnoreMe     int     `gorm:"-"` // 忽略本字段
}
```

### 结构体标记（tags）

使用结构体声明模型时，标记（tags）是可选项。gorm支持以下标记:

![](D:\myfile\GO\gin框架学习\pic\gorm标记1.jpg)

![](D:\myfile\GO\gin框架学习\pic\gorm标记2.jpg)

## 主键、表名、列名的约定

### 主键（Primary Key）

GORM 默认会使用名为ID的字段作为表的主键。

```go
type User struct {
  ID   string // 名为`ID`的字段会默认作为表的主键
  Name string
}

// 使用`AnimalID`作为主键
type Animal struct {
  AnimalID int64 `gorm:"primary_key"`
  Name     string
  Age      int64
}
```

### 表名（Table Name）

表名默认就是结构体名称的复数，例如：

```go
type User struct {} // 默认表名是 `users`

// 将 User 的表名设置为 `profiles`
func (User) TableName() string {
  return "profiles"
}
//也可以通通过用户属性为其生成相应的表名
func (u User) TableName() string {
  if u.Role == "admin" {
    return "admin_users"
  } else {
    return "users"
  }
}

// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
db.SingularTable(true)
```

也可以通过`Table()`指定表名：

```go
// 使用User结构体创建名为`deleted_users`的表
db.Table("deleted_users").CreateTable(&User{})

var deleted_users []User
db.Table("deleted_users").Find(&deleted_users)
//// SELECT * FROM deleted_users;

db.Table("deleted_users").Where("name = ?", "jinzhu").Delete()
//// DELETE FROM deleted_users WHERE name = 'jinzhu';
```

GORM还支持更改默认表名称规则：

```go
gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
  return "prefix_" + defaultTableName;
}
```

### 列名（Column Name）

列名由字段名称进行下划线分割来生成

```go
type User struct {
  ID        uint      // column name is `id`
  Name      string    // column name is `name`
  Birthday  time.Time // column name is `birthday`
  CreatedAt time.Time // column name is `created_at`
}
```

可以使用结构体tag指定列名：

```go
type Animal struct {
  AnimalId    int64     `gorm:"column:beast_id"`         // set column name to `beast_id`
  Birthday    time.Time `gorm:"column:day_of_the_beast"` // set column name to `day_of_the_beast`
  Age         int64     `gorm:"column:age_of_the_beast"` // set column name to `age_of_the_beast`
}
```

### 时间戳跟踪

#### CreatedAt

如果模型有 `CreatedAt`字段，该字段的值将会是初次创建记录的时间。

```go
db.Create(&user) // `CreatedAt`将会是当前时间

// 可以使用`Update`方法来改变`CreateAt`的值
db.Model(&user).Update("CreatedAt", time.Now())
```

#### UpdatedAt

如果模型有`UpdatedAt`字段，该字段的值将会是每次更新记录的时间。

```go
db.Save(&user) // `UpdatedAt`将会是当前时间

db.Model(&user).Update("name", "jinzhu") // `UpdatedAt`将会是当前时间
```

#### DeletedAt

如果模型有`DeletedAt`字段，调用`Delete`删除该记录时，将会设置`DeletedAt`字段为当前时间，而不是直接将记录从数据库中删除。



# CRUD

CRUD通常指数据库的增删改查操作，本文详细介绍了如何使用GORM实现创建、查询、更新和删除操作。

本文中的`db`变量为`*gorm.DB`对象，例如：

我们将连接数据库的配置放到dao/层下的mysql.go中

```go
package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)
//连接
func InitMysql() (err error) {
	dsn := "root:000000@tcp(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	fmt.Println("数据库连接成功")
	return DB.DB().Ping()
}
//关闭连接
func Close() {
	DB.Close()
}

```

## 创建

### 创建记录

首先定义模型：

在models层建立users.go

```go
package models

type User struct {
	ID   int64
	Name string `gorm:"default:'老王'"`
	Age  int64
}
//指定其表名为user_profile
func (u User) TableName() string {
	return "user_profile"
}

```

使用使用`NewRecord()`查询主键是否存在，主键为空使用`Create()`创建记录：

```go
user := User{Name: "q1mi", Age: 18}

db.NewRecord(user) // 主键为空返回`true`
db.Create(&user)   // 创建user
db.NewRecord(user) // 创建`user`后返回`false`
```

在main.go进行建表操作

```go
package main

import (
	"fmt"
	"grom_crud/dao"
	"grom_crud/models"
)

func main() {
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	fmt.Println("进行数据库创建表当中......")
	dao.DB.AutoMigrate(&models.User{})
	//具体代码操作
	user := models.User{Name: "小心", Age: 23}
	////判断该表的主键是否为空，如果是为空的话就返回true
	a := dao.DB.NewRecord(user)
	fmt.Println("该主键为空吗?true/false  ", a)
	//创建user
	dao.DB.Create(&user)
	dao.DB.NewRecord(user) //创建user后返回false

}
```



### 默认值

可以通过 tag 定义字段的默认值，比如：

```go
type User struct {
  ID   int64
  Name string `gorm:"default:'老王'"`
  Age  int64
}
```

**注意：**通过tag定义字段的默认值，在创建记录时候生成的 SQL 语句会排除没有值或值为 零值 的字段。 在将记录插入到数据库后，Gorm会从数据库加载那些字段的默认值。

举个例子：

```go
var user = User{Name: "", Age: 99}
db.Create(&user)
```

此时在表中创建的数据中Name字段就默认为"老王"

上面代码实际执行的SQL语句是`INSERT INTO users("age") values('99');`，排除了零值字段`Name`，而在数据库中这一条数据会使用设置的默认值`小王子`作为Name字段的值。

**注意：**所有字段的零值, 比如`0`, `""`,`false`或者其它`零值`，都不会保存到数据库内，但会使用他们的默认值。 如果你想避免这种情况，可以考虑使用指针或实现 `Scanner/Valuer`接口，比如：

#### 使用指针方式实现零值存入数据库

models层

```go
package models

type User struct {
	ID   int64
	Name *string `gorm:"default:'老王'"`
	Age  int64
}

func (u User) TableName() string {
	return "user_profile"
}

```

main.go

```go
package main

import (
	"fmt"
	"grom_crud/dao"
	"grom_crud/models"
)

//var (
//	db = dao.DB
//)

func main() {
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	fmt.Println("进行数据库创建表当中......")
	dao.DB.AutoMigrate(&models.User{})
	user := models.User{Name: new(string), Age: 18}
	dao.DB.Create(&user)

}
```

查看数据库可知新增的数据Name为0值

#### 使用Scanner/Valuer接口方式实现零值存入数据库

```go
// 使用 Scanner/Valuer
type User struct {
	ID int64
	Name sql.NullString `gorm:"default:'老王'"` // sql.NullString 实现了Scanner/Valuer接口
	Age  int64
}
user := User{Name: sql.NullString{"", true}, Age:18}
dao.DB.Create(&user)  // 此时数据库中该条记录name字段的值就是''
```

### 扩展创建选项

例如`PostgreSQL`数据库中可以使用下面的方式实现合并插入, 有则更新, 无则插入。

```go
// 为Instert语句添加扩展SQL选项
db.Set("gorm:insert_option", "ON CONFLICT").Create(&product)
// INSERT INTO products (name, code) VALUES ("name", "code") ON CONFLICT;
```



## 查询操作

### 1.前期准备

创建一个student（包含id,name,sex,age,score）的结构体，并通过数据迁移的形式在数据库中创建一个名为stu_profile的数据表

```go

package models

type Student struct {
	Id    uint `gorm:"primary key"`
	Name  string
	Sex   string
	Age   int64
	Score int
}

func (stu Student) TableName() string {
	return "stu_profile"
}



```

```go
package main

import (
	"fmt"
	"grom_crud/dao"
	"grom_crud/models"
)

//var (
//	db = dao.DB
//)

func main() {
	err := dao.InitMysql()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	fmt.Println("进行数据库创建表当中......")
	//创建数据库表结构
	dao.DB.AutoMigrate(&models.Student{})

	stu1 := models.Student{3, "小花", "男", 35, 90}
	dao.DB.Create(&stu1)

```

### 2.常用简单查询

**dao.DB.First(&object)**

```go
var stu models.Student
	dao.DB.First(&stu)//获得以主键为顺序的第一条记录
//相当于select * from stu_profile order by primary key;
```

如果设置的主键为数字，则可以将主键作为条件进行对象查找。当主键是字符串时则需要额外注意避免sql注入。

**dao.DB.First(&object)**

```go
var stu models.Student
	dao.DB.First(&stu,10)
//相当于select * from stu_profile where Id=10
类似的还有
dao.DB.First(&stu,10)
//select * from stu_profile where Id =10
dao.DB.First(&stus,[]int{1,2,3})
//select * from stu_profile where Id in (1,2,3)
//此外，如果主键是一个字符串(比如说是uuid),可以写成这样
dao.DB.First(&stu,"id=?","1c64513f-f2b8-409f-ac47-e8c062e3465a")
//相当于select * from stu_profile where Id="1c64513f-f2b8-409f-ac47-e8c062e3465a"
```

当目标对象有一个主键值，主键将会作为一个查询条件{Id:2}

```go
var stu = models.Student{Id:2}
dao.DB.First(&stu)
//相当于select * from stu_profile where Id = 2;
var res models.Student
db.Model(models.Student{Id:2}).First(&res)
//相当于select * from stu_profile were Id = 2;
```



**dao.DB.Take(&object)**

```go
var stu models.Student
	dao.DB.Take(&stu)//获取表中的一条记录，没有特别指定的顺序
//相当于select *from stu_profile limit 1;
```

**dao.DB.Last(&object)**

```go
var stu models.Student
dao.DB.Last(&stu)
```

注意:First和Last方法查询的出来的数据的顺序是以主键来参考的，如果定义的结构体字段中没有主键的话就使用结构体的第一个字段作为顺序参考

Note:如果使用gorm提供的特定变量类型比如gorm.DeletedAt,查出来的结果将会不同

```go
type User struct {
  ID           string `gorm:"primarykey;size:16"`
  Name         string `gorm:"size:24"`
  DeletedAt    gorm.DeletedAt `gorm:"index"`
}

var user = User{ID: 15}
db.First(&user)
//  SELECT * FROM `users` WHERE `users`.`id` = '15' AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1

```

查询所有对象操作

**dao.DB.Find(&objects)**

```go
var stus []models.Student
	dao.DB.Find(&stus)
	fmt.Println(stus)
//结果如下
                          
[{1 小明 男 24 78} {2 小红 女 25 80} {3 小花 男 35 90}]

additonal:
var stus []models.Student
	res := dao.DB.Find(&stus)
	//fmt.Println(res)
	fmt.Println(res.RowsAffected) //返回记录的总数。，相当于len(stus)
	fmt.Println(res.Error)        //返回错误
```

### 3.Where条件进行查询

```go
//获得符合条件的第一条记录（按照主键顺序）
var stu models.Student
	dao.DB.Where("Name=?", "小明").First(&stu)
	fmt.Println(stu)
//获得除了小红以外的所有记录
var stus []models.Student
	dao.DB.Where("name<>?","小红").Find(&stus)
	//Select * from stus Where Name <> '小红';

//IN
	var stus []models.Student
= dao.DB.Where("name IN ?", []string{"小明", "小红"}).Find(&stus)
//select * from stu_profile where name in('小明','小红')
//Like(模糊查询)
	dao.DB.Where("name LIKE ?", "%小明%").Find(&stus)
	//select * from stu_profile where name='%小明%'
//And
	dao.DB.Where("name = ? AND age >= ? ", "小明", "24").Find(&stus)
	//select * from stu_profile where name = '小明' and age>=24

// Time
db.Where("updated_at > ?", lastWeek).Find(&users)
// SELECT * FROM users WHERE updated_at > '2000-01-01 00:00:00';

//官网补充
// BETWEEN
db.Where("created_at BETWEEN ? AND ?", lastWeek, today).Find(&users)
// SELECT * FROM users WHERE created_at BETWEEN '2000-01-01 00:00:00' AND '2000-01-0
```

### 4.结构体&map 条件

```go
//struct
dao.DB.Where(&models.Student{Name: "小明", Age: 24}).First(&stu)
//相当于select * from stu_proofile where name = '小明' and age = 24 oeder by id limit 1;
//Map
dao.DB.Where(map[string]interface{}{"name": "小明", "age": 20}).Find(&stus)

//Slice
	dao.DB.Where([]int64{1, 2, 3}).Find(&stus)
```

注意：当你用结构体进行查询的时候，Gorm将只能和一个非0值字段进行查询。也意味着你的字段值为 0,' ',false或者其他非0值将不会被用于查询条件使用

```go
dao.DB.Where(&models.Student{Name: "小明",Age :24}).Find(&stus)
//相当于select * from stu_profile where name="小明"
```

如果你要将0值作为查询条件的话，你可以使用一个map,它的所包含的所有键值对都将作为查询条件

```go
dao.DB.Where(map[string]interface{}{"Name":"小明",”Age“:0}).Find(&stus)
```

当你用结构体进行查询的时候，你可以从结构体中指定一些特别的值去使用它并作为查询条件在相关的字段名或者数据库名例如

```go
dao.DB.Where(&models.Student{Name: "小明"}, "Name", "Age").Find(&stus)
//相当于select * frm stu_profile where name="小明" and age=0

dao.DB.Where(&models.Student{Name: "小明"},  "Age").Find(&stus)
//select * from stu_profile where age =0;
```

### 5.内联条件查询

查询条件可以被用作内敛方法像First和Find他们的功能类似于Where

```go
var stu models.Student
	dao.DB.First(&stu, "Id= ?", "1")
//select * from stu_profile where id='1';

//普通的SQL
	dao.DB.Find(&stu, "name=?", "小红")
//select * from stu_profile where name ="小红";

var stus []models.Student
	dao.DB.Find(&stus, "name <> ? AND age > ?", "小红", 25)
//相当于select * from stu_profile where name <> "小红"and age >25;表示查找name不等于小红以及age >35的对象

//struct
	dao.DB.Find(&stus, models.Student{Age: 25})
//相当于select * from stu_profile where age = 25;

//Map
	dao.DB.Find(&stus, map[string]interface{}{"age": 25})
```

### 6.Not条件

```go
//Not条件查询
	dao.DB.Not("Name", "小红").First(&stu)
//select * from stu_profile where name<> "小红" limit 1;
//Not In
dao.DB.Not("name", []string{"小红", "小明"}).Find(&stus)
//select * from stu_profile where name NOT in ("小红", "小明");
	dao.DB.Not([]int64{1, 3}).First(&stu)
//相当于select * from stu_profile where id Not IN(1,2)

	dao.DB.Not("name=?", "小红").First(&stu)
//相当于select * from stu_profile where not(name="小红")
//struct
	dao.DB.Not(models.Student{Name: "小明"}).First(&stu)
//相当于select * from stu_profile where name <>"小明"

```

### 7.or条件

```go
	dao.DB.Where("name= ?", "小红").Or("name=?", "小明").Find(&stus)
	fmt.Println(stus)
//select * from stu_profile where name='小红' or name='小明'

//struct 
	dao.DB.Where("name='小花").Or(models.Student{Name: "小红", Age: 25}).Find(&stus)
//select * from stu_profile where name ='小花'or(name='小红'AND age=25)

//Map
	dao.DB.Where("name='小花'").Or(map[string]interface{}{"name": "小红", "age": 25}).Find(&stus)
//select * from stu_profile where name='小花'or(name='小红'And age = 25)

//查询具体的字段
	dao.DB.Select("name", "age").Find(&stus)
//select name,age from stu_profile;

	dao.DB.Select([]string{"name", "age"}).Find(&stus)
	//select name,age from stu_profile;
	
	dao.DB.Table("stu_prfile").Select("COALESCE", 42).Rows()
//select COALESCE(age,'42') from stu_profile
```

### 8.Order

```go
//按照指定顺序进行查询操作
	//按照age降序的顺序进行查询
	dao.DB.Order("age desc,name").Find(&stus)
	//select * from stu_profile order by age desc,name;
	//多个顺序进行查找
	dao.DB.Order("age desc").Order("name").Find(&stus)
	//select * from stu_profile order by age desc,name;

dao.DB.Clauses(clause.OrderBy{
  Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
}).Find(&stus{})
// SELECT * FROM users ORDER BY FIELD(id,1,2,3)
```

### 9.limit&offset

Limit 查找指定记录最大数量去查询

Offset在执行记录之前跳过所指定的记录数

```go
var stus1 []models.Student
	var stus2 []models.Student
	//从表里查询三条记录
	dao.DB.Limit(3).Find(&stus)
	fmt.Println(stus)
	dao.DB.Limit(5).Find(&stus1).Limit(-1).Find(&stus2)
	//selcet * from stu_profile limit 5;(stus1)
	//selcet * from stu_profile ;(stus2)
	fmt.Println(stus1)
	fmt.Println(stus2)
	fmt.Println("\n")

	dao.DB.Offset(3).Find(&stus)
	//select * from stu_profile offset 3;
	fmt.Println(stus)
	dao.DB.Limit(10).Offset(5).Find(&stus)
	// SELECT * FROM stu_prfile  OFFSET 5 LIMIT 10;
	fmt.Println(stus)
	dao.DB.Offset(10).Find(&stus1).Offset(-1).Find(&stus2)
//select * from stu_profile offset 10;(stus1)
//select * from stu_profile ;(stus2)
```

 





# gorm的应用实战

清单小项目实现对数据库的增删查改的代码

main.go

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"net/http"
)

// 初始化一个全局的变量DB
var (
	DB *gorm.DB
)

// Todo
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:123456@tcp(192.168.20.10:3306)/bubble?charset=utf8mb4&parseTime=True"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	fmt.Println("数据库连接成功!")
	return DB.DB().Ping()

}

func main() {
	//创建数据库
	//sql：CREATE DATABASE bubble;
	//连接数据库
	err := initMySQL()
	if err != nil {
		panic(err) //数据库都连接不上的话就没有必要往下走了
	}
	//延迟注册关闭
	defer DB.Close() //程序退出关闭数据库连接
	//模型绑定
	DB.AutoMigrate(&Todo{}) //todos
	r := gin.Default()
	//告诉gin框架模板文件引用的静态文件去哪里找,要使用绝对路径
	r.Static("/css", "D:/myfile/GO/project/qindanprojects/myproject/bubble/static/css")
	r.Static("/js", "D:/myfile/GO/project/qindanprojects/myproject/bubble/static/js")
	r.Static("/fonts", "D:/myfile/GO/project/qindanprojects/myproject/bubble/static/fonts")

	//告诉gin框架去哪里找模板
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	//v1
	v1Group := r.Group("v1")
	{
		//代办事项
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			//前端页面填写待办事项，点击提交，会发送请求到这里
			//1.从请求中把数据拿出来
			var todo Todo
			c.BindJSON(&todo)
			//2.存入数据库
			//err =DB.Create(&todo).Error
			//if err != nil {
			//
			//}
			//简写成下面的代码
			//3.返回响应
			if err = DB.Create(&todo).Error; err != nil { //此代码就是将json中的数据存入到数据库中
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
				//在公司的写法
				//c.JSON(http.StatusOK,gin.H{
				//	"code":2000,
				//	"msg":"success",
				//	"data":todo,
				//})
			}

		})
		//查看所有的代办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			//查询todo这个表里的所有数据
			var todoList []Todo
			if err = DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		//查看某一个代办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})

		//修改某一个待办事项
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
				return
			}
			var todo Todo
			if err = DB.Where("id=?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": "无效id"})
				return
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		//删除某一个待办事项
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
				return
			}
			if err = DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})

	}

	r.Run(":9090")
}

```

