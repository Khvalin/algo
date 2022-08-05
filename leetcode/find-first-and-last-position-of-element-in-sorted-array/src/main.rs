struct Solution;

impl Solution {
    pub fn search_range(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let  r;

        let l = nums.partition_point(|&x| x < target);
        if l == nums.len() || nums[l] != target {
            return vec![-1, -1];
        } else {
            r = nums.partition_point(|&x| x <= target) -1;
        }

        vec![l as i32, r as i32]
    }
}

fn main() {
    let res = Solution::search_range(vec![5, 7, 7, 8, 8, 10], 8);

    println!("{:?}", res);
}
