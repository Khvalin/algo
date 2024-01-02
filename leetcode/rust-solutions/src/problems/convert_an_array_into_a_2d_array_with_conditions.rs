struct Solution;

impl Solution {
    pub fn find_matrix(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut res = vec![];
        let mut count = [0; 201];

        for n in &nums {
            let c = count[*n as usize];
            if c >= res.len() {
                res.push(vec![]);
            }
            res[c].push(*n);

            count[*n as usize] += 1;
        }

        res
    }
}
