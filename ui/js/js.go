/*
 * go4api - an api testing tool written in Go
 * Created by: Ping Zhu 2018
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 *
 */

package js

var Js = `
function mergeTcResults(setUpTcResults, normalTcResults, tearDownTcResults) {
	var c = setUpTcResults.concat(normalTcResults);
	var dest = c.concat(tearDownTcResults);

	return dest
}
`