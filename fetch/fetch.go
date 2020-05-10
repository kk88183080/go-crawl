package fetch

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// 10微秒请求一次
var ratelimit = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	// 延时请求
	//<-ratelimit

	//生成client 参数为默认
	client := &http.Client{}
	//log.Println("fetch请求地址", url)
	request, e := http.NewRequest(http.MethodGet, url, nil)

	if e != nil {
		panic(e)
	}

	// 设置cookies
	timestap := time.Now().Unix()
	cookie1 := &http.Cookie{Name: "Hm_lpvt_2c8ad67df9e787ad29dbd54ee608f5d2", Value: strconv.FormatInt(timestap, 10), HttpOnly: false}
	request.AddCookie(cookie1)

	cookie2 := &http.Cookie{Name: "FSSBBIl1UgzbN7NO", Value: "5o33UhOmILHRozGj9p.Ka1PmstJbJZkNpZupmw4BzIhAxsesFCK5nN6Cu7EF.6DbPqZSO1OQN7scchQuYI190Ea", HttpOnly: false}
	request.AddCookie(cookie2)

	// 设置头信息解决返回418的问题
	request.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.14; rv:75.0) Gecko/20100101 Firefox/75.0")
	request.Header.Add("Referer", "https://album.zhenai.com/u/1045778053")

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
	// 延时关闭
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("fetch Error status code:%d, %s ", resp.StatusCode, request.URL)
		fmt.Println()
	}

	bodyReader := bufio.NewReader(resp.Body)
	encode := DeterminEncode(bodyReader)

	reader := transform.NewReader(bodyReader, encode.NewDecoder())

	return ioutil.ReadAll(reader)
}

/**
解决编码问题
*/
func DeterminEncode(b *bufio.Reader) encoding.Encoding {
	rs, e := b.Peek(1024)
	if e != nil {
		//panic(e)
		return unicode.UTF8
	}

	e2, _, _ := charset.DetermineEncoding(rs, "")
	return e2
}
