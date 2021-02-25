/**
 * @param {number[][]} mat
 * @return {number[][]}
 */
var diagonalSort = function(mat) {
    
    const  sortDiagonal = (i, j) => {
    
    let diagonal = []

    // store the current diagonal 
    while (i < n && j < m) {
        diagonal.push(mat[i][j]);
        ++i;
        ++j;    
    }
    diagonal = diagonal.sort((a,b) => a-b)

    // push the sorted values 
    // back into the matrix
    while (i > 0 && j > 0) {
        --i;
        --j;
        mat[i][j] = diagonal.pop();    
    }
}
    
    
  const n = mat.length 
  const m = mat[0].length
    
 // sort all diagonals 
    // in the lower left corner
    for (let i = 0; i < n; ++i) {
        sortDiagonal(i, 0);       
    } 
    // sort all diagonals 
    // in the upper right corner
    for (let j = 0; j < m; ++j) {
        sortDiagonal(0, j);  
    } 
    return mat;
};
