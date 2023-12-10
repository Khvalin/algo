struct Solution;

impl Solution {
    pub fn transpose(matrix: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        if matrix.len() == 0 {
            return matrix;
        }
        
        let mut res = vec![vec![]; matrix[0].len()];
        
        for i in 0..matrix.len() {
            for j in 0..matrix[i].len() {
                res[j][i] = matrix[i][j];
            }
        }
        
        res
    }
}

fn main() {
    println!("Hello, world!");
}
