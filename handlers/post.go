package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "All posts")
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength          // 获取请求实体长度
	body := make([]byte, len)       // 创建存放请求实体的字节切片
	r.Body.Read(body)               // 调用 Read 方法读取请求实体并将返回内容存放到雌雄同体创建的字节切片
	io.WriteString(w, string(body)) // 将请求实体作为响应实体返回
}

func EditPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func EditPost2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.Form.Get("id")
	fmt.Println("post id:", id)
	fmt.Println("form data:", r.PostForm)

	r.ParseMultipartForm(1024)
	fmt.Println("post file:", r.MultipartForm)

	io.WriteString(w, "表单提交成功")
}

type Post struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

// curl.exe -X POST -id '{\"title\":\"test\",\"content\":\"hello\"}' "http://localhost:8080/post/add2" -H "Content-Type:application/json"
func AddPost2(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength    // 获取请求实体长度
	body := make([]byte, len) // 创建存放实体的字节切片
	r.Body.Read(body)         // 调用 Read 方法读取请求实体并将返回内容存放到上面创建的字节切片
	post := Post{}
	json.Unmarshal(body, &post)   // 对读取的 JSON 数据进行解析
	fmt.Fprintf(w, "%#v\n", post) // 格式化输出结果
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024 * 1000)      // 最大支持 1024 KB，即 1M
	name := r.MultipartForm.Value["name"]  // 文件名
	image := r.MultipartForm.File["image"] // 图片文件

	fmt.Println("图片上传成功：", name[0])

	file, err := image[0].Open()
	if err == nil {
		data, err := ioutil.ReadAll(file) // 读取二进制文件字节流
		if err == nil {
			fmt.Fprintln(w, string(data)) // 将读取的字节信息输出
			// 将文件存储到项目根目录下的 images 子目录
			// 从上传文件中读取文件名并获取文件后缀
			names := strings.Split(image[0].Filename, ".")
			suffix := names[len(names)-1]
			// 将上传文件名字段值和源文件后缀拼接出新的文件名
			filename := name[0] + "." + suffix
			// 创建这个文件
			newFile, _ := os.Create("images/" + filename)
			defer newFile.Close()
			// 将上传文件的二进制字节信息写入新建的文件
			size, err := newFile.Write(data)
			if err == nil {
				fmt.Fprintf(w, "图片上传成功，图片大小：%d 字节\n", size/1000)
			}
		}
	}
	if err != nil {
		fmt.Fprintln(w, err)
	}
}
