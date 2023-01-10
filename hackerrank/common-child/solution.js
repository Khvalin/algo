'use strict';

const fs = require('fs');

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString = '';
let currentLine = 0;

process.stdin.on('data', function(inputStdin) {
    inputString += inputStdin;
});

process.stdin.on('end', function() {
    inputString = inputString.split('\n');

    main();
});

function readLine() {
    return inputString[currentLine++];
}

/*
 * Complete the 'commonChild' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING s1
 *  2. STRING s2
 */

function commonChild(s1, s2) {
    let memo = new Array(s1.length + 1)
    for (let i = 0; i <= s1.length;i++) {
        memo[i] = memo[i] || new Array(s2.length + 1)
        for (let j = 0; j <= s2.length; j++) {
            if (i === 0 || j === 0) {
                
                memo[i][j] = 0
                continue
            }
            let r = Math.max(memo[i - 1][j], memo[i][j-1])
            if (s1[i-1] === s2[j-1]) {
               r = 1 + memo[i-1][j-1]
            }
            memo[i][j] = r
        }
    }

    return memo[s1.length][s2.length]
}

function main() {
    const ws = fs.createWriteStream(process.env.OUTPUT_PATH);

    const s1 = readLine();

    const s2 = readLine();

    const result = commonChild(s1, s2);

    ws.write(result + '\n');

    ws.end();
}
