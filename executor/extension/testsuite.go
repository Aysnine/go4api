/*
 * go4api - an api testing tool written in Go
 * Created by: Ping Zhu 2018
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
	"github.com/Aysnine/go4api/lib/extension/testsuite"
	"github.com/Aysnine/go4api/lib/testcase"
)

func GetTsFilePaths() []string {
	filePathSlice := strings.Split(cmd.Opt.Testsuite, ",")

	return filePathSlice
}

func InitFullTsTcSlice(filePaths []string) []*testcase.TestCaseDataInfo {
	// filePathSlice := GetTsFilePaths()

	fullTsTcSlice := testsuite.InitFullTsTcSlice(filePaths)

	return fullTsTcSlice
}
