struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn min_operations(nums: Vec<i32>) -> i32 {
        let counts = nums.into_iter().fold(HashMap::new(), |mut counts, n| {
            *counts.entry(n).or_insert(0) += 1;
            counts
        });

        let res = counts.values().into_iter().fold(Some(0), |res, c| {
            if res.is_none() {
                return res;
            }
            let mut t = 0;
            let mut n = *c;

            if n >= 3 {
                t += n / 3;
                n -= (n / 3) * 3;
            }

            if t > 0 && n == 1 {
                t -= 1;
                n += 3;
            }

            if n >= 2 {
                t += n / 2;
                n -= (n / 2) * 2;
            }

            if n > 0 {
                return None;
            }

            Some(t + res.unwrap())
        });

        if res.is_none() {
            return -1;
        }

        res.unwrap()
    }
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn min_operations() {
        assert_eq!(Solution::min_operations(vec![2, 3, 3, 2, 2, 4, 2, 3, 4]), 4);
        assert_eq!(
            Solution::min_operations(vec![
                14, 12, 14, 14, 12, 14, 14, 12, 12, 12, 12, 14, 14, 12, 14, 14, 14, 12, 12
            ]),
            7
        );
    }
}
