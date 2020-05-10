package main

import (
	"log"
	"strconv"
	"time"
)

func main() {
	log.Println(time.Now().Unix())
	log.Println(strconv.FormatInt(time.Now().Unix(), 10))
}
