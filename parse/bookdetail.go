package parse

import (
	"../engine"
	"../model"
	"fmt"
	"regexp"
	"strconv"
)

var testStr = `<div id="info" class="">



    
    
  
    <span>
      <span class="pl"> 作者</span>:
        
            
            <a class="" href="/search/%E5%BC%97%E6%9C%97%E7%B4%A2%E7%93%A6%E2%80%A2%E8%82%96%E8%8E%B1">[美] 弗朗索瓦•肖莱</a>
    </span><br>

    
    
  
    <span class="pl">出版社:</span> 人民邮电出版社<br>

    
    
  
    <span class="pl">出品方:</span>&nbsp;<a href="https://book.douban.com/series/47356?brand=1">图灵教育</a><br>

    
    
  

    
    
  
    <span class="pl">原作名:</span> Deep Learning with Python<br>

    
    
  
    <span>
      <span class="pl"> 译者</span>:
        
            
            <a class="" href="/search/%E5%BC%A0%E4%BA%AE">张亮</a>
    </span><br>

    
    
  
    <span class="pl">出版年:</span> 2018-8<br>

    
    
  
    <span class="pl">页数:</span> 292<br>

    
    
  
    <span class="pl">定价:</span> 119.00元<br>

    
    
  
    <span class="pl">装帧:</span> 平装<br>

    
    
  
    <span class="pl">丛书:</span>&nbsp;<a href="https://book.douban.com/series/660">图灵程序设计丛书</a><br>

    
    
  
    
      
      <span class="pl">ISBN:</span> 9787115488763<br>


</div>

<span class="pl">作者:</span>&nbsp;
        <a href="https://book.douban.com/author/1039386/">
                [哥伦比亚]
            加西亚·马尔克斯</a>
    <br>
`

var bookname_reg = regexp.MustCompile(`<h1>[\d\D]*?<span property="v:itemreviewed">([^<]+)</span>[\d\D]*?<div class="clear"></div>[\d\D]*?</h1>`)
var author_reg = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
var author_reg_blank = regexp.MustCompile(`<span class="pl">作者:</span>&nbsp;[\d\D]*?<a.*?>([^<]+)</a>`)
var publicer_reg = regexp.MustCompile(`<span class="pl">出版社:</span> ([^<]+)<br/>`)
var pages_reg = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br/>`)
var price_reg = regexp.MustCompile(`<span class="pl">定价:</span> ([^<]+)<br/>`)
var score_reg = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
var info_reg = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)

func ParseDetailContent(content []byte, bookname string) engine.ParseResult {
	//
	fmt.Printf("%s\n", content)

	result := engine.ParseResult{}
	bookdetai := model.Bookdetai{}
	if bookname == "" {
		bookdetai.Bookname = parseDetailItemVal(bookname_reg, content)
	} else {
		// 从列表页面获取数据
		bookdetai.Bookname = bookname
	}

	bookdetai.Author = parseDetailItemVal(author_reg, content)
	if bookdetai.Author == "" {
		bookdetai.Author = parseDetailItemVal(author_reg_blank, content)
	}

	bookdetai.Publicer = parseDetailItemVal(publicer_reg, content)
	pages, _ := strconv.Atoi(parseDetailItemVal(pages_reg, content))
	bookdetai.Pages = pages
	bookdetai.Price = parseDetailItemVal(price_reg, content)
	bookdetai.Score = parseDetailItemVal(score_reg, content)
	bookdetai.Info = parseDetailItemVal(info_reg, content)

	//log.Println(bookdetai)
	result.Items = []interface{}{bookdetai}

	return result
}

func parseDetailItemVal(reg *regexp.Regexp, content []byte) string {
	submatch := reg.FindSubmatch(content)

	if len(submatch) >= 2 {
		return string(submatch[1])
	}

	return ""
}
