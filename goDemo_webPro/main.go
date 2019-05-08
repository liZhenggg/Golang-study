// goDemo_webPro project main.go
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(">> sayHelloName()")

	//解析url传递的参数，对于POST则解析响应包的主体（request body）
	r.ParseForm()

	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	// fmt.Println("r.Form:", r.Form)
	// fmt.Println("path:", r.URL.Path)
	// fmt.Println("scheme:", r.URL.Scheme)
	// fmt.Println("r.Form[\"url_long\"]:", r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}

	//这个写入到w的是输出到客户端的
	fmt.Fprintf(w, "Hello World!")
}

func login(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(">> login()")

	//解析url传递的参数，对于POST则解析响应包的主体（request body）
	// Handler里面是不会自动解析form的
	// 若使用r.Form['xx']方式获取参数，必须显式的调用r.ParseForm()
	r.ParseForm()

	//获取请求方法
	// fmt.Println("method:", r.Method)

	//判断请求方式
	if r.Method == "GET" {
		t, _ := template.ParseFiles("pages/login.gtpl")
		t.Execute(w, nil)
	} else {

		//// 获取请求参数方式1：
		// 使用r.Form['xx']方式获取参数，必须显式的调用r.ParseForm()
		// r.Form里面包含了所有请求的参数，URL中的query-string字段和POST冲突时，会保存成一个slice，里面存储了多个值
		// username := r.Form["username"]
		// password := r.Form["password"]

		//// 获取请求参数方式2：
		// Request本身也提供了FormValue()函数来获取用户提交的参数。
		// 如r.Form["username"]也可写成r.FormValue("username")。
		// 调用r.FormValue 时会自动调用r.ParseForm，所以不必提前调用。
		//r.FormValue只会返回同名参数中的第一个，若参数不存在则返回空字符串。
		username := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Println("username:", username)
		fmt.Println("password:", password)

		if username == "admin" && password == "123123" {
			t, _ := template.ParseFiles("pages/index.gtpl")
			t.Execute(w, nil)
		} else {
			log.Fatalln("用户名或密码错误！")
		}

		//// request.Form是一个url.Values类型
		v := url.Values{}
		v.Set("name", "Ava")
		// v.Add("friend", "Jess")
		// v.Add("friend", "Sarah")
		// v.Add("friend", "Zoe")
		// // v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
		// fmt.Println(v.Get("name"))
		// fmt.Println(v.Get("friend"))
		// fmt.Println(v["friend"])
	}
}

type sysUser struct {
	nickname    string
	name_CN     string
	name_EN     string
	age         int
	gender      int
	idcard      string
	email       string
	nationality int
	hobby       []int
	birthday    time.Time
}

func userInfoEdit(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(">> userInfoEdit()")

	r.ParseForm()

	nickname := r.Form["nickname"]
	name_CN := r.Form["name_CN"]
	name_EN := r.Form["name_EN"]
	age := r.Form["age"]

	genderSlice := []string{"1", "0", "2"}

	var gender string
	// var gender = make([]string, 3)
	for _, v := range genderSlice {
		if v == r.Form.Get("gender") {
			fmt.Println("gender fetched:", v)
			gender = v
			// gender := append(gender, v)
		}
	}
	fmt.Println("gender:", gender)

	// idcard := r.Form["idcard"]
	// email := r.Form["email"]
	// nationality := r.Form["nationality"]
	// hobby := r.Form["hobby"]
	// birthday := r.Form["birthday"]

	if len(nickname[0]) == 0 {
		fmt.Println("昵称不能为空")
	}
	if len(name_CN[0]) == 0 {
		fmt.Println("中文名不能为空")
	}
	if len(name_EN[0]) == 0 {
		fmt.Println("英文名不能为空")
	}
	if len(age[0]) == 0 {
		fmt.Println("年龄不能为空")
	}
	// if len(gender[0]) == 0 {
	// 	fmt.Println("性别不能为空")
	// }
	// if len(idcard[0]) == 0 {
	// 	fmt.Println("身份证号不能为空")
	// }
	// if len(email[0]) == 0 {
	// 	fmt.Println("电子邮件不能为空")
	// }
	// if len(nationality[0]) == 0 {
	// 	fmt.Println("国家不能为空")
	// }
	// if len(hobby[0]) == 0 {
	// 	fmt.Println("爱好不能为空")
	// }
	// if len(birthday[0]) == 0 {
	// 	fmt.Println("生日不能为空")
	// }
}

func main() {
	// http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/userInfoEdit", userInfoEdit)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
