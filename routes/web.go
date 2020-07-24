package routes

import (
	"github.com/jaykof/goblog/handlers"
	"net/http"
)

// 定义一个 WebRoute 结构体用于存放单个路由
type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// 声明 WebRoutes 切片存放所有 Web 路由
type WebRoutes []WebRoute

// 定义所有 Web 路由
var webRoutes = WebRoutes{
	WebRoute{
		"Home",
		"GET",
		"/",
		handlers.Home,
	},
	WebRoute{
		"Posts",
		"GET",
		"/posts",
		handlers.GetPosts,
	},
	WebRoute{
		"User",
		"GET",
		"/user/{id}",
		handlers.GetUser,
	},
	WebRoute{
		"NewPost",
		"POST",
		"/post/add",
		handlers.AddPost,
	},
	WebRoute{
		"NewPost2",
		"POST",
		"/post/add2",
		handlers.AddPost2,
	},
	WebRoute{
		"UpdatePost",
		"POST",
		"/post/edit",
		handlers.EditPost,
	},
	WebRoute{
		"UpdatePost2",
		"POST",
		"/post/edit2",
		handlers.EditPost2,
	},
	WebRoute{
		"UploadImage",
		"POST",
		"/image/upload",
		handlers.UploadImage,
	},
	WebRoute{
		"ApiError",
		"GET",
		"/error",
		handlers.Error,
	},
	WebRoute{
		"Redirect",
		"GET",
		"/redirect",
		handlers.Redirect,
	},
	WebRoute{
		"SetCookie",
		"GET",
		"/setcookies",
		handlers.SetCookie,
	},
	WebRoute{
		"GetCookie",
		"GET",
		"/getcookies",
		handlers.GetCookie,
	},
	WebRoute{
		"SetMessage",
		"GET",
		"/set_welcome_message",
		handlers.SetWelcomeMessage,
	},
	WebRoute{
		"GetMessage",
		"GET",
		"/get_welcome_message",
		handlers.GetWelcomeMessage,
	},
}
