package curl

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"strings"
	//"net/url"
)
//get请求可以直接http.Get方法，非常简单。
func HttpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		println(err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		println(err.Error())
	}

	//fmt.Println(string(body))
	return string(body)
}
//Tips：使用这个方法的话，第二个参数要设置成”application/x-www-form-urlencoded”，否则post参数无法传递。
func HttpPost(url,data string) string {
	resp, err := http.Post(url,"application/x-www-form-urlencoded",strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	// handle error
		println(err.Error())
	}

	//fmt.Println(string(body))
	return string(body)
}

func HttpPostForm(url string,value map[string][]string) string {
    //resp, err := http.PostForm("http://www.01happy.com/demo/accept.php",
    //    url.Values{"key": {"Value"}, "id": {"123"}})
	resp, err := http.PostForm("http://www.01happy.com/demo/accept.php",
	value)

	if err != nil {
	// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	// handle error
	}

	//fmt.Println(string(body))
	return string(body)

}
//有时需要在请求的时候设置头参数、cookie之类的数据，就可以使用http.Do方法。
//同上面的post请求，必须要设定Content-Type为application/x-www-form-urlencoded，post参数才可正常传递。
//
//如果要发起head请求可以直接使用http client的head方法，比较简单，这里就不再说明。
//
//完整代码示例文件下载：golang_http_client发起get和post代码示例
func HttpDo(action,url,value string) string {
	client := &http.Client{}

	//req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
	req, err := http.NewRequest(action, url, strings.NewReader(value))
	if err != nil {
	// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", value)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	// handle error
	}

    //fmt.Println(string(body))
	return string(body)
}