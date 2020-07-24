package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Greeting struct {
	Message string `json:"message"`
}

func Home(w http.ResponseWriter, r *http.Request) {
	//io.WriteString(w, "Welcome to my blog site")
	//w.Write([]byte("欢迎访问我的个人网站😘💋"))
	/*html := `<html>
		<head>
			<title>我的个人网站</title>
		</head>
		<body>
			<h1>欢迎访问我的个人网站👀🐲</h1>
		</body>
	</html>`
	w.Write([]byte(html))*/
	// 返回 JSON 格式数据
	greeting := Greeting{
		"欢迎访问我的个人网站👀🐲",
	}
	message, _ := json.Marshal(greeting)
	w.Header().Set("Content-Type", "application/json")
	w.Write(message)
}

func Error(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
	fmt.Fprintln(w, "认证后才能访问该接口")
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	// 设置一个 301 重定向
	w.Header().Set("Location", "https://xueyuanjun.com")
	w.WriteHeader(301)
}

func SetCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "username",
		Value:    url.QueryEscape("我的名字111"),
		HttpOnly: true,
		Expires:  time.Now().AddDate(0, 0, 1), // Cookie 有效期设置为1天
	}
	c2 := http.Cookie{
		Name:     "website",
		Value:    "https://www.163.com",
		HttpOnly: true,
		MaxAge:   1000, // Cookie 有效期设置为1000s
	}
	//w.Header().Add("Set-Cookie", c1.String())
	//w.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
	fmt.Fprintln(w, "通过 HTTP 响应头发送 Cooike 信息")
}

func GetCookie(w http.ResponseWriter, r *http.Request) {
	c1, err := r.Cookie("username")
	if err != nil {
		fmt.Fprintln(w, "名为 username 的 Cookie不存在")
		return
	}
	username, _ := url.QueryUnescape(c1.Value)
	c2, err := r.Cookie("website")
	if err != nil {
		fmt.Fprintln(w, "名为 website 的 Cookie 不存在")
		return
	}
	website := c2.Value
	fmt.Fprintf(w, "从用户请求中读取的 Cookie: {username: %s, website: %s}\n", username, website)
}

func SetWelcomeMessage(w http.ResponseWriter, r *http.Request) {
	msg := "欢迎访问我的网站👏🤟"
	cookie := http.Cookie{
		Name:  "welcome_message",
		Value: base64.URLEncoding.EncodeToString([]byte(msg)),
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/get_welcome_message", 302)
}

func GetWelcomeMessage(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("welcome_message")
	if err != nil {
		fmt.Fprintln(w, "没有在 Cookie 中找到欢迎消息")
	} else {
		delCookie := http.Cookie{
			Name:   "welcome_message",
			MaxAge: -1,
		}
		http.SetCookie(w, &delCookie)
		msg, _ := base64.URLEncoding.DecodeString(cookie.Value)
		fmt.Fprintln(w, string(msg))
	}
}
