package model

import "strconv"

/**
详情信息
*/
type Bookdetai struct {
	// 书名
	Bookname string
	// 作者
	Author string
	// 出版社
	Publicer string
	// 页数
	Pages int
	// 价格
	Price string
	// 评分
	Score string
	// 内容简介
	Info string
}

func (b Bookdetai) String() string {
	return "书名:" + b.Bookname + ",\n" + "作者:" + b.Author + ",\n" + "出版社:" + b.Publicer + ",\n" + "页数:" + strconv.Itoa(b.Pages) + ",\n" + "价格:" + b.Price + ",\n" + "评分:" + b.Score + ",\n" + "内容简介:" + b.Info + ",\n"
}
