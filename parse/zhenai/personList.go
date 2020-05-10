package zhenai

import (
	"../../engine"
	"log"
	"regexp"
)

/*
`<div class="list-item">
<div class="photo">
<a href="http://album.zhenai.com/u/1468651805" target="_blank">
<img src="https://photo.zastatic.com/images/photo/367163/1468651805/10458977715177158.png?scrop=1&amp;crop=1&amp;w=140&amp;h=140&amp;cpos=north" alt="92Sun先生"></a><
/div> <div class="content"><table><tbody><tr><th>
<a href="http://album.zhenai.com/u/1468651805" target="_blank">92Sun先生</a></th></tr> <tr><
td width="180"><span class="grayL">性别：</span>男士</td>
<td><span class="grayL">居住地：</span>北京</td></tr>
<tr><td width="180"><span class="grayL">年龄：</span>27</td> <!---->
<td><span class="grayL">月&nbsp;&nbsp;&nbsp;薪：</span>8001-12000元</td></tr>
<tr><td width="180"><span class="grayL">婚况：</span>未婚</td>
<td width="180"><span class="grayL">身&nbsp;&nbsp;&nbsp;高：</span>175</td></tr></tbody>
</table>
<div class="introduce">92年的，月份比较小所以显示的是26岁，没啥写的，你对我一分的好，我会对你万分的好。</div>
</div> <div class="item-btn">打招呼</div></div>`
*/

var reg = regexp.MustCompile(`<img src="([^"]+)" alt="([^"]+)"></a></div> <div class="content"><table><tbody><tr><th><a href="([^"]+)" target="_blank">([^<]+)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>([^<]+)</td>`)

func ParsePersonList(content []byte, city string) engine.ParseResult {
	result := engine.ParseResult{}
	log.Println(string(content))
	submatch := reg.FindAllSubmatch(content, -1)
	if len(submatch) == 0 {
		log.Println(city, "没有匹配到数据")
	}
	for _, rs := range submatch {
		log.Println("person list name:%s, url:%s", string(rs[1]), string(rs[2]), string(rs[3]), string(rs[4]), string(rs[5]))
		result.Items = append(result.Items, rs[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(rs[3]),
			ParseFunc: func(bytes []byte) engine.ParseResult {
				return ParsePersonDetail(bytes, string(rs[1]), string(rs[5]))
			},
		})

	}

	return result
}
