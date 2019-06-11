// goDemo_webPro project main.go
package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"

	"go/main/utils"
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
	mobile      string
	nationality int
	hobby       []int
	birthday    time.Time
}

func userInfoEdit(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(">> userInfoEdit()")

	r.ParseForm()

	//// 防XSS攻击，处理特殊字符
	// nickname := r.Form["nickname"]
	nickname := template.HTMLEscapeString(r.Form.Get("nickname"))
	fmt.Println("nickname:", nickname)
	//输出信息到客户端
	template.HTMLEscape(w, []byte(r.Form.Get("nickname")))

	// t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	// err = t.ExecuteTemplate(out, "T", "<scirpt>alert('you have been pwned')</scirpt>")

	//// 判断中文方式一：正则表达式
	if m, _ := regexp.MatchString("^\\p{Han}+$", r.Form.Get("name_CN")); !m {
		fmt.Println("中文名输入有误！")
	} else {
		fmt.Println("中文名输入正确！")
	}
	name_CN := r.Form["name_CN"]
	fmt.Println("name_CN:", name_CN)

	// 判断中文方式2：使用 unicode 包提供的 func Is(rangeTab *RangeTable, r rune) bool
	// todo...

	//// 判断英文
	if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("name_EN")); !m {
		fmt.Println("英文名输入有误！")
	} else {
		fmt.Println("英文名输入正确！")
	}
	name_EN := r.Form["name_EN"]
	fmt.Println("name_EN:", name_EN)

	//// 判断单选框
	genderSlice := []string{"1", "0", "2"}
	var gender string
	for _, v := range genderSlice {
		if v == r.Form.Get("gender") {
			// fmt.Println("gender fetched:", v)
			gender = v
		}
	}
	if gender == "" {
		fmt.Println("gender 未选择!")
	} else {
		fmt.Println("gender:", gender)
	}

	//// 判断数字--方式1
	getint, err := strconv.Atoi(r.Form.Get("age"))
	if err != nil {
		//转数字出错，输入不是正确的数字，getint将被赋值为0
		fmt.Println("输入年龄格式不正确")
	} else if getint > 100 || getint < 1 {
		fmt.Println("年龄范围有误")
		// return false
	} else {
		fmt.Println("年龄输入正确！")
	}

	//判断数字方式2--正则表达式
	if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
		fmt.Println("年龄输入有误!！")
	} else {
		fmt.Println("年龄输入正确！！")
	}
	age := r.Form["age"]
	fmt.Println("age:", age)

	//// 判断Email
	if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
		fmt.Println("email输入有误！")
	} else {
		fmt.Println("email输入正确！")
	}
	email := r.Form.Get("email")
	fmt.Println("email:", email)

	//// 判断手机号码
	if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
		fmt.Println("手机号码输入有误！")
	} else {
		fmt.Println("手机号码输入正确！")
	}
	mobile := r.Form.Get("mobile")
	fmt.Println("mobile:", mobile)

	//// 验证下拉菜单
	var nationality string
	nati_slice := []string{"0", "1", "2", "3", "4"}
	// nati_slice := []int{0, 1, 2, 3, 4}
	v := r.Form.Get("nationality")
	for _, item := range nati_slice {
		if item == v {
			nationality = v
		}
	}

	if nationality == "0" {
		fmt.Println("nationality 未选择")
	} else {
		fmt.Println("nationality:", nationality)
	}

	////验证复选框
	hobby_slice := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	hobby := sliceUtils.Slice_diff(r.Form["hobby"], hobby_slice)
	if hobby == nil {
		fmt.Println("爱好为空")
	} else {
		fmt.Println("hobby:", hobby)
	}

	////验证身份证号码
	if m, _ := regexp.MatchString(`^(\d{15})$`, r.Form.Get("idcard")); m {
		//验证15位身份证，15位的是全部数字
		fmt.Println("是有效的15位身份证号码！")
	} else if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("idcard")); m {
		//验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
		fmt.Println("是有效的18位身份证号码！")
	} else {
		fmt.Println("身份证号码有误！")
	}
	idcard := r.Form.Get("idcard")
	fmt.Println("idcard:", idcard)

	////验证日期
	birthday, err := time.Parse("2006-01-02 15:04:05", r.Form.Get("birthday"))
	if err == nil {
		fmt.Println("birthday:", birthday)
	} else {
		fmt.Println("日期转换错误！")
		fmt.Println("err:", err)
	}

	////
	// if len(nickname[0]) == 0 {
	// 	fmt.Println("昵称不能为空")
	// }
	// if len(name_CN[0]) == 0 {
	// 	fmt.Println("中文名不能为空")
	// }
	// if len(name_EN[0]) == 0 {
	// 	fmt.Println("英文名不能为空")
	// }
	// if len(age[0]) == 0 {
	// 	fmt.Println("年龄不能为空")
	// }
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

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("pages/upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		//获取文件名 uploadfile
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		//指定存放文件位置
		f, err := os.OpenFile("./headImages/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) // 此处假设当前目录下已存在test目录
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
	}
}

//上传文件
func postFile(filename string, targetUrl string) error {
	fmt.Println(">>>>postFile  ")
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// 关键操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error wirting to buffer")
		return err
	}

	// 打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	// iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil
}

func main() {

	target_url := "http://localhost:9000/upload"
	filename := "./test.jpg"
	postFile(filename, target_url)

	// http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	http.HandleFunc("/userInfoEdit", userInfoEdit)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
