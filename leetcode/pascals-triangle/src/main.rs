struct Solution;

impl Solution {
    pub fn generate(num_rows: i32) -> Vec<Vec<i32>> {
        let mut res: Vec<Vec<i32>> = vec![vec![1]];

        for i in 1..num_rows as usize {
            let mut a = vec![1];
            for j in 1..res[i - 1].len() {
                a.push(res[i - 1][j] + res[i - 1][j - 1]);
            }
            a.push(1);

            res.push(a);
        }

        return res;
    }
}

fn main() {
    let r = Solution::generate(2);
    print!("{:?}", r);
}
