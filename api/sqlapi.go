/*
 * go4api - a api testing tool written in Go
 * Created by: Ping Zhu 2018
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 *
 */

package api

import (
    "fmt"
    "strings" 
    "encoding/json"

    "go4api/sql"
    "go4api/lib/testcase"

    gjson "github.com/tidwall/gjson"
)
// Note: for each SetUp, TesrDown, it may have more than one Command (including sql)
// for each Command, it may have more than one assertion
func (tcDataStore *TcDataStore) RunTcSetUp () (string, [][]*testcase.TestMessage) {
    tcData := tcDataStore.TcData
    cmdGroup := tcData.TestCase.SetUp()

    finalResults, finalTestMessages := tcDataStore.CommandGroup("setUp", cmdGroup)

    return finalResults, finalTestMessages
}

func (tcDataStore *TcDataStore) RunTcTearDown () (string, [][]*testcase.TestMessage) {
    tcData := tcDataStore.TcData
    cmdGroup := tcData.TestCase.TearDown()

    finalResults, finalTestMessages := tcDataStore.CommandGroup("tearDown", cmdGroup)

    return finalResults, finalTestMessages
}

func (tcDataStore *TcDataStore) CommandGroup (section string, cmdGroup []*testcase.CommandDetails) (string, [][]*testcase.TestMessage) {
    finalResults := "Success"
    var cmdsResults []bool
    var finalTestMessages = [][]*testcase.TestMessage{}
    //
    cmdGroupJsonB, _ := json.Marshal(cmdGroup)
    cmdGroupJson := string(cmdGroupJsonB)

    for i, _ := range cmdGroup {
        cmdType := gjson.Get(cmdGroupJson, fmt.Sprint(i) + "." + "cmdType")

        switch strings.ToLower(cmdType.String()) {
            case "sql":
                cmdStrPath := "TestCase." + tcDataStore.TcData.TestCase.TcName() + "." + section + "." + fmt.Sprint(i) + ".cmd"
                tcDataStore.RenderTcVariables(cmdStrPath)
                tcDataStore.EvaluateTcBuiltinFunctions(cmdStrPath)
                //
                cmdStr := gjson.Get(cmdGroupJson, fmt.Sprint(i) + "." + "cmd")
                fmt.Println("cmdStr: ", cmdStr)
                rowsCount, _, rowsData, sqlExecStatus := RunSql(cmdStr.String())

                if sqlExecStatus == "SqlSuccess" {
                    path := "TestCase." + tcDataStore.TcData.TestCase.TcName() + "." + section + "." + fmt.Sprint(i) + ".cmdResponse"
                    tcDataStore.RenderTcVariables(path)
                    tcDataStore.EvaluateTcBuiltinFunctions(path)
                    //
                    cmdExpResp := gjson.Get(cmdGroupJson, fmt.Sprint(i) + "." + "cmdResponse").Map()

                    singleCmdResults, testMessages := tcDataStore.CompareRespGroup(cmdExpResp, rowsCount, rowsData)

                    // HandleSingleCommandResults for out
                    if singleCmdResults == true {
                        cmdDetails := cmdGroup[i]
                        tcDataStore.HandleCmdResultsForOut(section, cmdDetails, i, rowsCount, rowsData)
                    }

                    cmdsResults = append(cmdsResults, singleCmdResults)
                    finalTestMessages = append(finalTestMessages, testMessages)
                } else {
                    cmdsResults = append(cmdsResults, false)
                }
            case "redis":
                fmt.Println("tbd, redis is not ready yet")
            default:
                fmt.Println("!! warning, command ", cmdType.String(), " can not be recognized.")
        }
    }

    for key := range cmdsResults {
        if cmdsResults[key] == false {
            finalResults = "Fail"
            break
        }
    }

    return finalResults, finalTestMessages
}

func (tcDataStore *TcDataStore) CompareRespGroup (cmdExpResp map[string]gjson.Result, 
        rowsCount int, rowsData []map[string]interface{}) (bool, []*testcase.TestMessage){
    //-----------
    singleCmdResults := true
    var testResults []bool
    var testMessages []*testcase.TestMessage

    for key, value := range cmdExpResp {
        cmdExpResp_sub := value.Value().(map[string]interface{})
        for assertionKey, expValueOrigin := range cmdExpResp_sub {
            
            actualValue := tcDataStore.GetResponseValue(key, rowsCount, rowsData)

            var expValue interface{}
            switch expValueOrigin.(type) {
                case float64, int64: 
                    expValue = expValueOrigin
                default:
                    expValue = tcDataStore.GetResponseValue(expValueOrigin.(string), rowsCount, rowsData)
            }
            
            testRes, msg := compareCommon("sql", key, assertionKey, actualValue, expValue)
            
            testMessages = append(testMessages, msg)
            testResults = append(testResults, testRes)
        }
    }

    for key := range testResults {
        if testResults[key] == false {
            singleCmdResults = false
            break
        }
    }

    return singleCmdResults, testMessages
}

func (tcDataStore *TcDataStore) HandleCmdResultsForOut (section string, cmdDetails *testcase.CommandDetails, i int, rowsCount int, rowsData []map[string]interface{}) {
    // write out session if has
    path := "TestCase." + tcDataStore.TcData.TestCase.TcName() + "." + section + "." + fmt.Sprint(i) + ".session"
    tcDataStore.RenderTcVariables(path)
    tcDataStore.EvaluateTcBuiltinFunctions(path)

    expTcSession := cmdDetails.Session
    tcDataStore.WriteSession(expTcSession, rowsCount, rowsData)

    // write out global variables if has
    path = "TestCase." + tcDataStore.TcData.TestCase.TcName() + "." + section + "." + fmt.Sprint(i) + ".outGlobalVariables"
    tcDataStore.RenderTcVariables(path)
    tcDataStore.EvaluateTcBuiltinFunctions(path)

    expOutGlobalVariables := cmdDetails.OutGlobalVariables
    tcDataStore.WriteOutGlobalVariables(expOutGlobalVariables, rowsCount, rowsData)

    // write out tc loca variables if has
    path = "TestCase." + tcDataStore.TcData.TestCase.TcName() + "." + section + "." + fmt.Sprint(i) + ".outLocalVariables"
    tcDataStore.RenderTcVariables(path)
    tcDataStore.EvaluateTcBuiltinFunctions(path)

    expOutLocalVariables := cmdDetails.OutLocalVariables
    tcDataStore.WriteOutGlobalVariables(expOutLocalVariables, rowsCount, rowsData)
}

func RunSql (stmt string) (int, []string, []map[string]interface{}, string) {
    // gsql.Run will return: <impacted rows : int>, <rows for select : [][]interface{}{}>, <sql status : string>
    // status: SqlSuccess, SqlFailed
    rowsCount, rowsHeaders, rowsData, sqlExecStatus := gsql.Run(stmt)

    return rowsCount, rowsHeaders, rowsData, sqlExecStatus
}


