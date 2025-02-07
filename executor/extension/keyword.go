/*
 * go4api - an api testing tool written in Go
 * Created by: Ping Zhu 2019.07
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 *
 */

package extension

import (
	// "fmt"
	"strings"

	"github.com/Aysnine/go4api/cmd"
	"github.com/Aysnine/go4api/lib/extension/keyword"
	"github.com/Aysnine/go4api/lib/testcase"
)

func GetKwFilePaths() []string {
	filePathSlice := strings.Split(cmd.Opt.KeyWord, ",")

	return filePathSlice
}

func InitFullKwTcSlice(filePaths []string) ([]*testcase.TestCaseDataInfo, []string) {
	// filePathSlice := GetTsFilePaths()

	fullKwTcSlice, fullKwJsSlice := keyword.InitFullKwTcSlice(filePaths)

	return fullKwTcSlice, fullKwJsSlice
}
