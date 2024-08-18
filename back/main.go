package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
)

func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}

func login(c *gin.Context) {
	// 登录信息
	username := c.PostForm("username")
	password := c.PostForm("password")
	// 判断是否为空
	if username == "" || password == "" {
		c.JSON(500, gin.H{
			"message": "登录信息不完整",
		})
		return
	}
	//连接到MYSQL
	usr := "huhaoyu:huhaoyu555@tcp(localhost:3306)/cugoj?charset=utf8"
	db, _ := sql.Open("mysql", usr)
	//查找账户密码是否存在
	rows, err := db.Query("select * from user where username=? and password=?", username, password)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "（数据库错误）" + err.Error(),
		})
		return
	} else {
		if rows.Next() {
			//生成token，有效时间1天
			nowTime := time.Now()
			expireTime := nowTime.Add(24 * time.Hour)
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": username,
				"password": password,
				"exp":      expireTime.Unix(),
			})
			tokenString, err := token.SignedString([]byte("secret"))
			if err != nil {
				c.JSON(500, gin.H{
					"message": "生成token失败" + err.Error(),
				})
				return
			}
			c.JSON(200, gin.H{
				"token": tokenString,
			})
		} else {
			c.JSON(500, gin.H{
				"message": "登录失败（账号或密码错误）",
			})
		}
	}
	//释放连接
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
}

func register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("mail")
	code := c.PostForm("code")
	//判断是否为空
	if username == "" || password == "" || email == "" {
		c.JSON(500, gin.H{
			"message": "注册信息不完整",
		})
		return
	}
	//判断邮箱验证码是否正确
	usr := "huhaoyu:huhaoyu555@tcp(localhost:3306)/cugoj?charset=utf8"
	db, _ := sql.Open("mysql", usr)
	rows, err := db.Query("select * from emailCode where email=? and code=? and valid_time>?",
		email, code, time.Now().Add(8*time.Hour))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "（数据库错误）" + err.Error(),
		})
		return
	} else {
		if !rows.Next() {
			c.JSON(500, gin.H{
				"message": "邮箱验证码错误",
			})
			return
		}
	}
	//判断用户名是否重复
	rows, err = db.Query("select * from user where username=?", username)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "（数据库错误）" + err.Error(),
		})
		return
	} else {
		if rows.Next() {
			c.JSON(500, gin.H{
				"message": "用户名重复",
			})
			return
		}
	}
	//生成token，有效时间1天
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
		"mails":    email,
		"exp":      expireTime.Unix(),
	})
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "生成token失败" + err.Error(),
		})
		return
	}
	//连接到MYSQL
	_, err = db.Exec("insert into user"+
		"(username,password,email,black_list,admin,register_time) values(?,?,?,?,?,?)",
		username, password, email, 0, 0, time.Now().Add(8*time.Hour))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "（数据库错误）" + err.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"token": tokenString,
		})
	}
	//释放连接
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
}

func mail(c *gin.Context) {
	//查询
	email := c.PostForm("mail")
	//生成邮箱验证码
	var arr = []int{1, 2, 3, 4, 5, 6}
	Seed := rand.NewSource(time.Now().UnixNano())
	MyRand := rand.New(Seed)
	MyRand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	var verifyCode string
	for i := 0; i < len(arr); i++ {
		verifyCode += strconv.Itoa(arr[i])
	}
	//发送邮件
	host := "smtp.qq.com"
	port := 465
	username := "1009422458@qq.com"
	password := "kjfsxlkeuqdpbfbe"
	d := gomail.NewDialer(
		host,
		port,
		username,
		password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	m := gomail.NewMessage()
	m.SetHeader("From", username)
	m.SetHeader("To", email)
	m.SetBody("text/plain", "CUG_Online_Judge的邮箱验证码为:"+verifyCode+"    （有效期2小时）")
	//发送邮件
	if err := d.DialAndSend(m); err != nil {
		c.JSON(500, gin.H{
			"message": "发送失败" + err.Error(),
		})
	}
	//把(key:username,value:邮箱的验证码)存到数据库里面
	usr := "huhaoyu:huhaoyu555@tcp(localhost:3306)/cugoj?charset=utf8"
	db, _ := sql.Open("mysql", usr)
	//转换为datatime(注意时区)
	_, err := db.Exec("insert into emailCode(email,code,valid_time) values(?,?,?)",
		email, verifyCode, time.Now().Add(2*time.Hour).Add(8*time.Hour))
	//返回信息
	if err != nil {
		c.JSON(500, gin.H{
			"message": "（数据库错误）" + err.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"message": "发送成功",
		})
	}
	//释放连接
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
}

func test() {
	// 获取本地所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("Error getting network interfaces:", err)
		return
	}

	// 遍历每个接口
	for _, i := range interfaces {
		// 获取接口的所有地址
		addrs, err := i.Addrs()
		if err != nil {
			fmt.Println("Error getting addresses for interface:", i.Name, " :", err)
			continue
		}

		// 遍历每个地址
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// 确保是IPv4地址且不是本地地址
			if ip.To4() != nil && !ip.IsLoopback() {
				fmt.Println("Interface:", i.Name, " IP Address:", ip.String())
			}
		}
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	//test()
	r := gin.New()
	r.Use(Cors())
	r.POST("/api/register", register)
	r.POST("/api/login", login)
	r.POST("/api/mail", mail)
	//输出当前的IP地址
	err := r.Run("172.19.231.116:8080")
	if err != nil {
		return
	}
}
