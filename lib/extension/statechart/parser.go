/*
 * go4api - an api testing tool written in Go
 * Created by: Ping Zhu 2019
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 *
 */

package statechart

import (
	"fmt"
	"os"
	"sync"

	// "strings"
	// "bufio"
	// "io"
	// "path/filepath"
	"encoding/json"

	"github.com/Aysnine/go4api/lib/testcase"
	"github.com/Aysnine/go4api/utils"
)

func InitFullScTcSlice(scfilePathSlice []string) []*testcase.TestCaseDataInfo {
	var fullScTcSlice []*testcase.TestCaseDataInfo
	// var fullKwJsPathSlice []string

	fmt.Println(scfilePathSlice)

	for i, _ := range scfilePathSlice {
		// scFileListTemp, _ := utils.WalkPath(scfilePathSlice[i], ".scxml")
		scFileListTemp, _ := utils.WalkPath(scfilePathSlice[i], ".xstate")

		for _, path := range scFileListTemp {
			// content := utils.GetContentFromFile(path)
			// XmlDecode(content)

			ConstructXstate(path)
		}
	}

	return fullScTcSlice
}

func ConstructXstate(xstateFile string) {
	var xstate State
	var transitions []*Transition

	jsonStr := utils.GetJsonFromFile(xstateFile)

	err := json.Unmarshal([]byte(jsonStr), &xstate)
	if err != nil {
		fmt.Println("!! Error, parse xstate into xstate failed: ", xstateFile, ". Cause: ", err)
		os.Exit(1)
	}

	// reset the state id
	xstate.SetStateIds()

	// here to use channel to avoid the global variable for transitions
	ch := make(chan *Transition)

	go func(ch chan *Transition) {
		defer close(ch)
		wg := &sync.WaitGroup{}

		wg.Add(1)
		xstate.GetStateTransitions(ch, wg)

		wg.Wait()
	}(ch)

	for v := range ch {
		transitions = append(transitions, v)
		kwStr := v.FromState + " " + v.Event + " " + v.ToState
		fmt.Println("111: ", kwStr)
	}

	fmt.Println("222: ", transitions)
}
