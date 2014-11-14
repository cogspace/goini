package wini

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/bmizerany/assert"
)

func Test1(t *testing.T) {

	filename := filepath.Join(testDataDir(), "ini_parser_testfile.ini")
	ini := New()
	err := ini.ParseFile(filename)
	assert.Equal(t, nil, err)

	v, ok := ini.Get("", "mid")
	assert.Equal(t, v, "ac9219aa5232c4e519ae5fcb4d77ae5b")
	assert.Equal(t, ok, true)

	v, ok = ini.Get("", "product")
	assert.Equal(t, v, "ppp")
	assert.Equal(t, ok, true)

	v, ok = ini.Get("", "combo")
	assert.Equal(t, v, "ccc")
	assert.Equal(t, ok, true)

	v, ok = ini.Get("", "aa")
	assert.Equal(t, v, "bb")
	assert.Equal(t, ok, true)

	v, ok = ini.Get("", "axxxa")
	assert.Equal(t, v, "")
	assert.Equal(t, ok, false)

	m, ok := ini.GetKvmap("")
	assert.Equal(t, len(m), 6)
	assert.Equal(t, ok, true)

	n, ok := ini.GetKvmap("n")
	assert.Equal(t, len(n), 0)
	assert.Equal(t, ok, false)

	sss, ok := ini.GetKvmap("sss")
	assert.Equal(t, len(sss), 2)
	assert.Equal(t, ok, true)
	v, ok = ini.Get("sss", "aa")
	assert.Equal(t, v, "bb")
	assert.Equal(t, ok, true)
	v, ok = ini.Get("sss", "appext")
	assert.Equal(t, v, "ab=cd")
	assert.Equal(t, ok, true)
}


func TestUft8(t *testing.T) {
	/*
		title=百度搜索_ipad2
		url=http://www.baidu.com/s?bs=ipad&f=8&rsv_bp=1&wd=ipad2&inputT=397
		url_md5=5844a75423cd3372e1997360bd110a25
		refer=http://www.google.com
		anchor_text= google
		ret_form = json
		 ret_start = 0
		 ret_limit =    50
		page_info   =  0,0,50,1,0,20
		local=0
		mid=c4ca4238a0b923820dcc509a6f75849b
		product=test
		combo=test
		version=1.0.0.1
		debug=1
		encoding=1

	*/

	filename := filepath.Join(testDataDir(), "utf8.ini")
	ini := New()
	err := ini.ParseFile(filename)
	assert.Equal(t, nil, err)

	v, ok := ini.Get("", "title")
	assert.Equal(t, v, "百度搜索_ipad2")
	assert.Equal(t, ok, true)

	v, ok = ini.Get("", "url_md5")
	assert.Equal(t, v, "5844a75423cd3372e1997360bd110a25")
	assert.Equal(t, ok, true)

	v, ok = ini.Get("", "ret_form")
	assert.Equal(t, v, "json")
	assert.Equal(t, ok, true)

	v, ok = ini.Get("", "ret_start")
	assert.Equal(t, v, "0")
	assert.Equal(t, ok, true)


	v, ok = ini.Get("", "ret_limit")
	assert.Equal(t, v, "50")
	assert.Equal(t, ok, true)



	v, ok = ini.Get("", "axxxa")
	assert.Equal(t, v, "")
	assert.Equal(t, ok, false)

	m, ok := ini.GetKvmap("")
	assert.Equal(t, len(m), 16)
	assert.Equal(t, ok, true)

	n, ok := ini.GetKvmap("n")
	assert.Equal(t, len(n), 0)
	assert.Equal(t, ok, false)
}

func TestErrorFormat(t *testing.T) {
		filename := filepath.Join(testDataDir(), "error.ini")
	ini := New()
	err := ini.ParseFile(filename)
	assert.NotEqual(t, nil, err)
}

func TestMemoryData1(t *testing.T) {
	//TODO
}


func testDataDir() string {
	var file string
	var ok bool
	if _, file, _, ok = runtime.Caller(0); ok {
		fmt.Printf("file=%v\n", file)
	}

	curdir := filepath.Dir(file)
	return filepath.Join(curdir, "test/data")
}
