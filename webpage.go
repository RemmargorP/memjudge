package myjudge

import (
	"io/ioutil"
)

type WebPage struct {
	Page string
	Data map[interface{}]interface{}
}

func getPage(name string) (*WebPage, error) {
	buf, err := ioutil.ReadFile(PublicDir + name)
	return &WebPage{Page: string(buf[:len(buf)]), Data: make(map[interface{}]interface{})}, err
}
