struct Solution;

impl Solution {
    pub fn length_of_lis(nums: Vec<i32>) -> i32 {
        let p_nums = &nums;
        let mut memo = p_nums.into_iter().map(|_| 1).collect::<Vec<_>>();

        let mut res = 0;
        if nums.len() > 0 {
            res = 1;
        }

        for i in 1..p_nums.len() {
            for j in 0..i {
                if p_nums[i] <= p_nums[j] {
                    continue;
                }

                memo[i] = memo[i].max(memo[j] + 1);
                res = res.max(memo[i]);
            }
        }

        res
    }
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn length_of_lis() {
        assert_eq!(Solution::length_of_lis(vec![10, 9, 2, 5, 3, 7, 101, 18]), 4);
        assert_eq!(Solution::length_of_lis(vec![7, 7, 7, 7, 7, 7, 7]), 1);
    }
}
