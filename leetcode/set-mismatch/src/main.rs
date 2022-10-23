impl Solution {
    pub fn find_error_nums(nums: Vec<i32>) -> Vec<i32> {
        let mut len = nums.len();
        let mut count = vec![0; len + 1];

        let mut res = vec![0, 0];
        for n in nums {
            count[n as usize] += 1
        }

        for i in 1..=len {
            if count[i] == 2 {
                res[0] = i as i32;
            }
            if count[i] == 0 {
                res[1] = i as i32;
            }
        }

        res
    }
}

fn main() {
    println!("Hello, world!");
}
