package main

import "C"
import (
	"encoding/json"
	"github.com/GopeedLab/gopeed/bind"
	"github.com/GopeedLab/gopeed/pkg/api"
	"github.com/GopeedLab/gopeed/pkg/api/model"
)

func main() {}

//export Create
func Create(cfg *C.char) *C.char {
	var config model.StartConfig
	if err := json.Unmarshal([]byte(C.GoString(cfg)), &config); err != nil {
		return C.CString(bind.BuildResult(err))
	}

	return C.CString(bind.Create(&config))
}

//export Invoke
func Invoke(index int, req *C.char) *C.char {
	var request api.Request
	if err := json.Unmarshal([]byte(C.GoString(req)), &request); err != nil {
		return C.CString(bind.BuildResult(err))
	}

	return C.CString(bind.Invoke(index, &request))
}
