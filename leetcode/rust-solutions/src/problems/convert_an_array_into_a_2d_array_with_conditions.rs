struct Solution;

impl Solution {
    pub fn find_matrix(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut res = Vec::with_capacity(nums.len());
        let mut count = [0; 201];

        for n in &nums {
            let c = count[*n as usize];
            if c >= res.len() {
                res.push(Vec::with_capacity(nums.len()));
            }
            res[c].push(*n);

            count[*n as usize] += 1;
        }

        res
    }
}
