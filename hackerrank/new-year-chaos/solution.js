'use strict';

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString = '';
let currentLine = 0;

process.stdin.on('data', function (inputStdin) {
    inputString += inputStdin;
});

process.stdin.on('end', function () {
    inputString = inputString.split('\n');

    main();
});

function readLine() {
    return inputString[currentLine++];
}

/*
 * Complete the 'minimumBribes' function below.
 *
 * The function accepts INTEGER_ARRAY q as parameter.
 */

function minimumBribes(q) {
    // Write your code here
    for (let i = 0; i < q.length; i++) {
        q[i]--
    }

    let res = 0
    for (let i = 0; i < q.length; i++) {
        let ind = 0
        for (let j = 0; j < q.length; j++) {
            if (i === q[j]) {
                ind = j
                break
            }
        }

        const c = i - ind

        if (c === 0) {
            continue
        }

        if (c > 2 || c < 0) {
            console.log('Too chaotic')
            return
        }
        while (i !== ind) {
            const k = ind + 1;
            [q[ind], q[k]] = [q[k], q[ind]]
            ind += 1
        }
        res += c
    }

    for (let i = 0; i < q.length; i++) {
        if (q[i] !== i) {
            console.log('Too chaotic')
            return
        }
    }

    console.log(res)
}

function main() {
    const t = parseInt(readLine().trim(), 10);

    for (let tItr = 0; tItr < t; tItr++) {
        const n = parseInt(readLine().trim(), 10);

        const q = readLine().replace(/\s+$/g, '').split(' ').map(qTemp => parseInt(qTemp, 10));

        minimumBribes(q);
    }
}

