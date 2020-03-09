package main

import (
	"log"
	"net/http"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main(){
	page := getPages()
}

func getPages() int {
	res, err := http.Get(baseURL)
	checkErr(err)
	checkRes(res)
	return 0
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkRes(res *http.Response){
	if res.StatusCode != 200 {
		log.Fatalln("Fail Status:", res.StatusCode)
	}
}