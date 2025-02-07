/*
 * go4api - an api testing tool written in Go
 * Created by: Ping Zhu 2018
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 *
 */

package gjs

import (
	"fmt"
	// "os"
	"path/filepath"
	"strings"

	// "github.com/Aysnine/go4api/cmd"
	"github.com/Aysnine/go4api/utils"

	"github.com/dop251/goja"
	"github.com/dop251/goja_nodejs/require"
)

var JsFunctions []GJsBasics

// trial code for js
func InitJsFunctions(jsFileList []string) {
	for i, _ := range jsFileList {
		srcBytes := utils.GetContentFromFile(jsFileList[i])
		src := string(srcBytes)

		p, err := goja.Compile("", src, false)
		if err != nil {
			panic(err)
		}

		jsFileName := strings.TrimSuffix(filepath.Base(jsFileList[i]), ".js")

		jsFunc := GJsBasics{
			JsSourceFilePath: jsFileList[i],
			JsSourceFileName: filepath.Base(jsFileList[i]),
			JsFunctionName:   jsFileName,
			JsProgram:        p,
		}

		JsFunctions = append(JsFunctions, jsFunc)
	}
}

// for testing
func CallJsFuncs(funcName string, funcParams interface{}) interface{} {
	// fmt.Println(JsFunctions)
	for ii, _ := range JsFunctions {
		for i := 1; i < 10; i++ {
			go RunProgram(JsFunctions[ii].JsProgram, funcParams)
		}
	}

	return 1
}

// for testing
func CallJsFunc(funcName string, funcParams interface{}) interface{} {
	// fmt.Println(JsFunctions)
	idx := -1
	var returnValue interface{}

	for i, _ := range JsFunctions {
		if JsFunctions[i].JsFunctionName == funcName {
			idx = i
			break
		}
	}

	if idx != -1 {
		returnValue = RunProgram(JsFunctions[idx].JsProgram, funcParams)
	} else {
		fmt.Println("! Error, no js function found")
	}

	return returnValue
}

func RunProgram(p *goja.Program, funcParams interface{}) interface{} {

	vm := goja.New()

	registry := new(require.Registry) // this can be shared by multiple runtimes
	registry.Enable(vm)

	vm.Set("funcParams", funcParams)

	v, err := vm.RunProgram(p)

	if err != nil {
		panic(err)
	}

	return v.Export()
}
