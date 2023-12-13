struct Solution;

impl Solution {
    pub fn num_special(mat: Vec<Vec<i32>>) -> i32 {
        let mut res = 0;

        for row in &mat {
            let mut col = 0;
            let mut sum = 0;
            for i in 0..(&row).len() {
                sum += &row[i];
                if row[i] == 1 {
                    col = i;
                }
            }

            if sum != 1 {
                continue;
            }

            for j in 0..(&mat[col]).len() {
                sum -= mat[col][j];
            }

            if sum == 0 {
                res += 1;
            }
        }

        res
    }
}
