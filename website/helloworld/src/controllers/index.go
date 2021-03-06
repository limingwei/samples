package controllers

import (
	"fmt"
	"net/http"
	"time"
)

type PageIndex struct {
	Application
}

func init() {
	App.RegisterController("index/", PageIndex{})
}

func (p *PageIndex) Before_() []map[string]string {
	return []map[string]string{
		map[string]string{
			"_FILTER": "getIndex",
			"_ALL":    "allow",
			"Index":   "deny",
		},
		map[string]string{
			"_FILTER":  "GetTest",
			"TestPage": "allow",
		},
		map[string]string{
			"_FILTER":  "GetTest1",
			"TestPage": "allow",
			"Index":    "allow",
		},
	}
}

func (p *PageIndex) After_() []map[string]string {
	return []map[string]string{
		map[string]string{
			"_FILTER":  "after",
			"_ALL":     "allow",
			"TestPage": "deny",
		},
	}
}

func (p *PageIndex) After_Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("After_Index")
}

func (p *PageIndex) Index(w http.ResponseWriter, r *http.Request) {
	p.SESSION["String"] = "String"
	p.SESSION["string"] = "string"
	p.SESSION["Int"] = 1
	p.SESSION["Map"] = map[string]string{
		"a": "b",
		"b": "c",
	}
	p.SESSION["Int64"] = time.Now().UnixNano()

	p.Body = func(a, b int) int {
		return a + b
	}
}

func (p *PageIndex) TestPage(w http.ResponseWriter, r *http.Request) {
	p.Document.Title = "测试页面"
	w.Write([]byte(fmt.Sprintf("%v", p.SESSION["String"])))
	w.Write([]byte(fmt.Sprintf("%v", p.SESSION["string"])))
	w.Write([]byte(fmt.Sprintf("%v", p.SESSION["Int"])))
	w.Write([]byte(fmt.Sprintf("%v", p.SESSION["Map"])))
	w.Write([]byte(fmt.Sprintf("%v", p.SESSION["Int64"])))
}

func (p *PageIndex) Filter_getIndex() {
	fmt.Println("GetIndex")
}

func (p *PageIndex) Filter_GetTest(r *http.Request) {
	fmt.Println("GetTest")
	fmt.Println(r.URL)
}

func (p *PageIndex) Filter_GetTest1() {
	fmt.Println("GetTest1")
}

func (p *PageIndex) Filter_after() {
	fmt.Println("Filter_after")
}
