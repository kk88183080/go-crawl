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
	"time"
)

// 10微秒请求一次
var ratelimit = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	// 延时请求
	<-ratelimit

	//生成client 参数为默认
	client := &http.Client{}

	request, e := http.NewRequest(http.MethodGet, url, nil)

	if e != nil {
		panic(e)
	}

	// 设置头信息解决返回418的问题
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36")

	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	// 延时关闭
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%d ", resp.StatusCode)
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
