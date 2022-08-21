/*
 * go4api - an api testing tool written in Go
 * Created by: Ping Zhu 2018
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 *
 */

package api

import (
	// "os"
	// "fmt"
	// "strings"
	// "encoding/json"

	"github.com/Aysnine/go4api/cmd"
	gsession "github.com/Aysnine/go4api/lib/session"
	"github.com/Aysnine/go4api/utils"
)

// data lookup sequence, latter override former one(s)
// config (json) -> env variables (key, value) -> session (key, value) -> data file (*_dt) / data file (inputs)

// Note: here may occur: fatal error: concurrent map iteration and map write, => need to fix
func (tcDataStore *TcDataStore) MergeTestData() map[string]interface{} {
	var finalMap = make(map[string]interface{})
	// 1
	envMap := utils.GetOsEnviron()
	for k, v := range envMap {
		finalMap[k] = v
	}

	// 2, options, cmdArgs
	for k, v := range cmd.CmdArgs {
		finalMap[k] = v
	}

	globalVariables := gsession.LoopGlobalVariables()
	for k, v := range globalVariables {
		finalMap[k] = v
	}

	// take care the test suite params
	tsName := tcDataStore.TcData.TestCase.TestSuite()
	tsSessionMap := gsession.LookupTcSession(tsName)
	for k, v := range tsSessionMap {
		finalMap[k] = v
	}

	// 3
	pTcName := tcDataStore.TcData.TestCase.ParentTestCase()
	pSssionMap := gsession.LookupTcSession(pTcName)
	for k, v := range pSssionMap {
		finalMap[k] = v
	}

	//
	tcName := tcDataStore.TcData.TestCase.TcName()
	sessionMap := gsession.LookupTcSession(tcName)
	for k, v := range sessionMap {
		finalMap[k] = v
	}

	//
	tcLocalVariables := tcDataStore.TcLocalVariables
	for k, v := range tcLocalVariables {
		finalMap[k] = v
	}

	finalMap["body"] = string(tcDataStore.HttpActualBody)

	return finalMap
}

// 4. data file (*_dt) / data file (inputs)
func ConvertCsvRowToMap(csvHeader []string, csvRow []interface{}) map[string]interface{} {
	csvMap := map[string]interface{}{}

	for i, item := range csvRow {
		csvMap[csvHeader[i]] = item
	}

	return csvMap
}
