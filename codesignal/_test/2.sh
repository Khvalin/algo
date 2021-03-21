#!/bin/bash
# Test for no cities

read -r -d '' serverScript << EOM

/* -------------- NodeJS server code Start ------------- */
const http = require('http')
const port = 80

function generateMap() {
  return [
    {
      "x": 10,
      "y": 10,
      "type": "street",
      "name": "Street_1",
      "length": 10,
      "is_busy": true
    },
    {
      "x": 0,
      "y": 0,
      "type": "street",
      "name": "street_1_1",
    },
    {
      "x": -1,
      "y": 10,
      "type": "street",
      "name": "street_2",
      "area": 400
    },
    {
      "x": -2,
      "y": 4,
      "type": "building",
      "name": "Building_1"
    }
  ];
}

let map = generateMap();
let mandatoryFields = ['x_from', 'x_to', 'y_from', 'y_to']

let borders = {
  "x": [-2, 10],
  "y": [-5, 15]
};

let K = 4;

const requestHandler = (request, response) => {
  const url = require('url');
  const querystring = require('querystring');

  var path = url.parse(request.url).pathname;

  if (path == '/maxResponsePatchSize') {
    response.end(JSON.stringify({"K": K}));
    return;
  }

  if (path == '/borders') {
    const bordersAnswer = {"x_min": borders.x[0], "x_max": borders.x[1],
                           "y_min": borders.y[0], "y_max": borders.y[1]};
    response.end(JSON.stringify(bordersAnswer));
    return;
  }

  if (path == '/map') {
    try {
      let query = querystring.parse(url.parse(request.url).query);
      const allFieldsThere = mandatoryFields.every(field => field in query);
      if (!allFieldsThere) {
        throw new Error('No field');
      }
      Object.keys(query).forEach(key => query[key] = parseInt(query[key]));

      let elems = map.filter(point =>
        (query.x_from <= point.x && point.x <= query.x_to &&
          query.y_from<= point.y && point.y <= query.y_to))

      if (elems.length > K) {
        response.end(JSON.stringify({"message": "Error: too many objects to return"}));
        return;
      }

      response.end(JSON.stringify(elems));
      return;
    }
    catch (error) {
      response.end(JSON.stringify({"message": 'Error: Invalid Input'}));
      return;
    }
  }
  response.end(JSON.stringify({"message": 'Error'}));
}

const server = http.createServer(requestHandler)

server.listen(port, (err) => {
  if (err) {
    return console.log('something bad happened', err)
  }

  console.log('server is listening on ' + port);
})
/* -------------- NodeJS server code End ------------- */

EOM


# Preparing SQL server
rm -rf /root/devops
mkdir /root/devops

# Preparing nodejs server
touch /root/devops/server.js
echo -e "$serverScript" >> /root/devops/server.js
ps aux | grep server.js | awk '{print $2}' | xargs kill &>/dev/null
nohup node /root/devops/server.js &>/dev/null &
sleep 1


