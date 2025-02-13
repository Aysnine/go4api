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
	// "fmt"

	gsession "github.com/Aysnine/go4api/lib/session"
)

func (tcDataStore *TcDataStore) WriteOutGlobalVariables(expOutGlobalVariables map[string]interface{}) {
	if expOutGlobalVariables != nil {
		for k, v := range expOutGlobalVariables {
			var value interface{}

			switch v.(type) {
			case string:
				value = tcDataStore.GetResponseValue(v.(string))
			case int, int64, float64:
				value = v
			}

			gsession.WriteGlobalVariables(k, value)
		}
	}
}

func (tcDataStore *TcDataStore) WriteOutTcLocalVariables(expOutLocalVariables map[string]interface{}) {
	if expOutLocalVariables != nil {
		for k, v := range expOutLocalVariables {
			var value interface{}

			switch v.(type) {
			case string:
				value = tcDataStore.GetResponseValue(v.(string))
			case int, int64, float64:
				value = v
			}

			tcDataStore.TcLocalVariables[k] = value
		}
	}
}
