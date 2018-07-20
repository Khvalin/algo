/**
 * @param {number} N
 * @return {number}
 */
var binaryGap = function(N) {
    const binary = N.toString(2);
    let pos1 = -1;
    let pos2 = -1;

    let maxGap = 0;

    for (let i =0; i < binary.length;i++) {
        const d = binary[i];

        if (d === '1'){
            if (pos1 >= 0){
                maxGap = Math.max(maxGap, i - pos1)
            }
            pos1 = i;
        }
    }

    return maxGap;
};

console.log(binaryGap(8))
console.log(binaryGap(6))
console.log(binaryGap(5))