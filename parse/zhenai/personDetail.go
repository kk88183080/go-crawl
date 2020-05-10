package zhenai

import (
	"../../engine"
	"../../model/zhenaiModel"
	"log"
	"regexp"
	"strings"
)

//([^<]+)
var reg_name = regexp.MustCompile(`<h1 class="nickName" data-v-5b109fc3>([^<]+)</h1>`)
var reg_id = regexp.MustCompile(`<div class="id" data-v-5b109fc3>ID：([^<]+)</div>`)

//var reg_more_item = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>([^|]+) | ([^|]+) | ([^|]+) | ([^|]+) | ([^|]+) | ([^|]+)</div>`)
var reg_more_item = regexp.MustCompile(`<div class="des f-cl" data-v-3c42fade>([^<]+)</div>`)
var reg_desc = regexp.MustCompile(`<div class="m-content-box m-des" data-v-8b1eac0c><span data-v-8b1eac0c>([^<]+)</span></div>`)

func ParsePersonDetail(content []byte, img string, sex string) engine.ParseResult {

	person := zhenaiModel.Person{}
	person.Photo = img
	person.Name = detailItemFind(reg_name, content)
	person.Id = detailItemFind(reg_id, content)
	person.Introduce = detailItemFind(reg_desc, content)
	person.Sex = sex
	detailItemMoreFind(reg_more_item, content, &person)

	result := engine.ParseResult{}
	result.Requests = nil
	result.Items = []interface{}{person}

	return result
}

func detailItemMoreFind(reg *regexp.Regexp, content []byte, person *zhenaiModel.Person) {
	rs := reg.FindSubmatch(content)
	log.Println(len(rs))
	if len(rs) >= 1 {
		log.Print(string(rs[1]))
		split := strings.Split(string(rs[1]), "|")
		if len(split) >= 6 {
			person.City = split[0]
			person.Age = split[1]
			person.Schoole = split[2]
			person.Status = split[3]
			person.Height = split[4]
			person.Money = split[5]
		}
	}

}

func detailItemFind(reg *regexp.Regexp, content []byte) string {
	rs := reg.FindSubmatch(content)
	if len(rs) >= 1 {
		log.Print(string(rs[1]))
		return string(rs[1])
	}

	return ""
}
