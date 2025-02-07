/*
 * go4api - an api testing tool written in Go
 * Created by: Ping Zhu 2018
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 *
 */

package ui

var Graphic = `<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
    <link href="style/go4api.css" rel="stylesheet" type="text/css"/>
    <script type="text/javascript" src="js/go4api.js"></script>
    <script type="text/javascript" src="js/results.js"></script>
    <script type="text/javascript" src="js/stats.js"></script>
    <script type="text/javascript" src="js/graphic.js"></script>
  <title>Go4Api Reports</title>
</head>
<body>
  <div class="container">
      <div class="head">
          <a href="https://github.com/zpsean/go4api" target="blank_" title="Go4Api Home Page"><img alt="Go4Api" src="style/logo.png"/></a>
      </div>
      <div class="main">
          <div class="skeleton">
              <div class="content">
                  <div class="sous-menu">
                      <div class="item "><a href="index.html">Overview</a></div>
                      <div class="item "><a id="details_link" href="details.html">Details</a></div>
                      <div class="item selected"><a id="graphic_link" href="graphic.html">Graphic</a></div>
                      <div class="item "><a id="mindex_link" href="mindex.html">MutationOverview</a></div>
                      <div class="item "><a id="mutation_link" href="mutation.html">Mutation</a></div>
                      <div class="item "><a id="fuzz_link" href="fuzz.html">Fuzz</a></div>

                      <script type="text/javascript">
                        // var timestamp = 1523957748602;
                        // var runStartHumanDate = new Date(timestamp).format("YYYY-MM-DD HH:mm:ss Z");
                        var runStartHumanDate = gStart.substr(0, 19)
                        var runDuration = (gEndUnixNano - gStartUnixNano) / 1000000000
                        document.writeln("<p class='sim_desc'>");
                        document.writeln("<b>" + "Started at " + runStartHumanDate + ", duration: " + runDuration + " seconds </b>");
                        document.writeln("</p>");
                      </script>
                  </div>

                  <div class="content-in">
                    <h1><span>> </span>Overview Information</h1>
                    <div class="article">


                      <svg width="1150" height="8000" id="mainSVG" version="1.1" xmlns="http://www.w3.org/2000/svg">
                          <line x1="0" y1="0" x2="0" y2="8000" style="stroke:rgb(99,99,99);stroke-width:2"/>
                          <line x1="0" y1="0" x2="1080" y2="0" style="stroke:rgb(99,99,99);stroke-width:2"/>
                          <line x1="1080" y1="0" x2="1080" y2="8000" style="stroke:rgb(99,99,99);stroke-width:1"/>
                          <line x1="0" y1="8000" x2="1080" y2="8000" style="stroke:rgb(99,99,99);stroke-width:2"/>
                      </svg>

                      <script type="text/javascript">
                              
                              var svgRoot = document.getElementById('mainSVG');

                              for (var key in circles)
                              {
                                  var index = indexOfResults(key)

                                  var parsed = null
                                  if (tcResults[index]) {
                                    if (tcResults[index].ActualBody) {
                                      try {
                                        parsed = JSON.parse(tcResults[index].ActualBody)
                                      } catch(error) {
                                        // ...
                                      }
                                    }
                                  }

                                  if (parsed) {
                                    tcResults[index].ActualBody = parsed
                                  }

                                  var strJons = JSON.stringify(tcResults[index], null, 4)

                                  var c = document.createElementNS('http://www.w3.org/2000/svg', 'circle');
                                  c.setAttribute('cx', circles[key][0]);
                                  c.setAttribute('cy', circles[key][1]);
                                  c.r.baseVal.value = circles[key][2];
                                  c.setAttribute('fill', circles[key][3]);

                                  svgRoot.appendChild(c);

                                  ;(function(_c, _key, _strJons) {
                                      _c.addEventListener('mouseenter', () => {
                                          var location = _c.getBoundingClientRect()
                                          var offsetleft = location.x
                                          var offsettop =  location.y
                                          console.log(offsetleft, offsettop, circles[_key][0], circles[_key][1])
                                          _c.setAttribute('fill', "blue");

                                          var modal = document.getElementById('modal')
                                          modal.style.display = "block"
                                          modal.style.left = offsetleft + circles[_key][2] + "px"
                                          modal.style.top = 190 + circles[_key][1] + "px"
                                          modal.innerText = _strJons
                                      })
                                      _c.addEventListener('mouseleave', () => {
                                          _c.setAttribute('fill', circles[_key][3]);

                                          // var modal = document.getElementById('modal')
                                          // modal.style.display = "None"
                                      })
                                  })(c, key, strJons)
                              }

                              for (var key in priorityLines)
                              {
                                  var line = document.createElementNS('http://www.w3.org/2000/svg', 'line');
                                          
                                  line.setAttribute('x1', priorityLines[key][0]);
                                  line.setAttribute('y1', priorityLines[key][1]);
                                  line.setAttribute('x2', priorityLines[key][2]);
                                  line.setAttribute('y2', priorityLines[key][3]);
                                  line.setAttribute('style', "stroke:rgb(99,99,99);stroke-width:1");

                                  svgRoot.appendChild(line);
                              }

                              for (var key in parentChildrenlines)
                              {
                                  var line = document.createElementNS('http://www.w3.org/2000/svg', 'line');
                                          
                                  line.setAttribute('x1', parentChildrenlines[key][0]);
                                  line.setAttribute('y1', parentChildrenlines[key][1]);
                                  line.setAttribute('x2', parentChildrenlines[key][2]);
                                  line.setAttribute('y2', parentChildrenlines[key][3]);
                                  line.setAttribute('style', "stroke:rgb(250,99,99);stroke-width:1");

                                  svgRoot.appendChild(line);
                              }

                              function indexOfResults(tcName){
                                for (var i = 0; i < tcResults.length; i++)
                                {
                                  if (tcResults[i].TcName == key)
                                  {
                                    return i
                                  }
                                }
                              }

                      </script>

                    </div>
                  </div>
              </div>
          </div>
      </div>
  </div>
  <div class="foot">
      <a href="https://github.com/zpsean/go4api" title="Go4Api Home Page"><img alt="Go4Api" src="style/logosmall.png"/></a>
  </div>

  <pre id="modal" style="display:none; border:1px solid #dcdcdc; position: absolute; left:0; top:0; background-color:#F5F5F5; width: 800px; white-space: pre-wrap; word-wrap:break-word;word-break:normal; opacity: 0.9"></pre>

  <script type="text/javascript">
    window.onclick = function() {
      var modal = document.getElementById('modal')
      modal.style.display = "None"
    }
  </script>

</body>
</html>
`
