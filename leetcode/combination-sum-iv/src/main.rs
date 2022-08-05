struct Solution;

impl Solution {
    pub fn combination_sum4(nums: Vec<i32>, target: i32) -> i32 {
        let mut arr: Vec<Option<i32>> = vec![None; 1001];
        fn solve(nums: &Vec<i32>, a: &mut Vec<Option<i32>>, t: i32) -> i32 {
            if t < 0 {
                return 0;
            }

            let ind = t as usize;

            if a[ind].is_some() {
                return a[ind].unwrap();
            }
            let mut res = 0;

            for n in nums {
                if *n <= t {
                    res += solve(nums, a, t - *n);
                }
            }

            a[ind] = Some(res);

            res
        }

        arr[0] = Some(1);
        solve(&nums, &mut arr, target)
    }
}

fn main() {
    println!("{}", Solution::combination_sum4(vec![1, 2, 3], 4));
    println!("{}", Solution::combination_sum4(vec![1, 2, 3], 14));
}
