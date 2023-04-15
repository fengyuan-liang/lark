package utils

import (
	"bytes"
	jsoniter "github.com/json-iterator/go"
	"lark/pkg/common/xlog"
	"strings"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(in interface{}) (str string, err error) {
	var (
		buf []byte
	)
	buf, err = json.Marshal(in)
	if err != nil {
		return
	}
	str = string(buf)
	return
}

func ObjToByte(in interface{}) (buf []byte, err error) {
	buf, err = json.Marshal(in)
	return
}

func ByteToObj(buf []byte, out interface{}) (err error) {
	dc := json.NewDecoder(bytes.NewReader(buf))
	dc.UseNumber()
	return dc.Decode(out)
}

func Unmarshal(in string, out interface{}) error {
	//return json.Unmarshal([]byte(in), out)
	dc := json.NewDecoder(strings.NewReader(in))
	dc.UseNumber()
	return dc.Decode(out)
}

func Copy(src interface{}, dst interface{}) (err error) {
	var (
		buf []byte
	)
	buf, err = json.Marshal(src)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(buf, dst); err != nil {
		return err
	}
	return
}

func ObjToMap(in interface{}) map[string]interface{} {
	var (
		maps map[string]interface{}
		buf  []byte
		err  error
	)
	if buf, err = json.Marshal(in); err != nil {
		//fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(buf))
		d.UseNumber()
		if err = d.Decode(&maps); err != nil {
			//fmt.Println(err)
		} else {
			for k, v := range maps {
				maps[k] = v
			}
		}
	}
	return maps
}

func ObjToJsonStr(obj interface{}) (str string) {
	str = ""
	var err error
	if obj == nil {
		xlog.Error("obj is nil")
		return
	}
	str, err = Marshal(obj)
	if err != nil {
		xlog.Error(err)
		return
	}
	return
}
