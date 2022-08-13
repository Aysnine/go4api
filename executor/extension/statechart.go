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
	"github.com/Aysnine/go4api/lib/extension/statechart"
	"github.com/Aysnine/go4api/lib/testcase"
)

func GetScFilePaths() []string {
	filePathSlice := strings.Split(cmd.Opt.StateChart, ",")

	return filePathSlice
}

func InitFullScTcSlice(filePaths []string) []*testcase.TestCaseDataInfo {
	// filePathSlice := GetTsFilePaths()

	fullKwTcSlice := statechart.InitFullScTcSlice(filePaths)

	return fullKwTcSlice
}
