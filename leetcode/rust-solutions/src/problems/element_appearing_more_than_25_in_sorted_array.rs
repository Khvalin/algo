struct Solution;

impl Solution {
    pub fn find_special_integer(arr: Vec<i32>) -> i32 {
        let n = arr.len();
        let candidates = vec![arr[n / 4], arr[n / 2], arr[3 * n / 4]];
        let target = n / 4;

        for c in candidates {
            let left = arr.partition_point(|&x| x < c) + 1;
            let right = arr.partition_point(|&x| x <= c);

            if right - left + 1 > target {
                return c;
            }
        }

        unreachable!()
    }
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_find_special_integer() {
        assert_eq!(Solution::find_special_integer(vec![1, 1]), 1);
        assert_eq!(
            Solution::find_special_integer(vec![1, 2, 2, 6, 6, 6, 6, 7, 10]),
            6
        );
    }
}
