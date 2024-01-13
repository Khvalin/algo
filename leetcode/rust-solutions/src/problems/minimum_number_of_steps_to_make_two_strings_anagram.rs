struct Solution;

impl Solution {
    pub fn min_steps(s: String, t: String) -> i32 {
        fn count_chars(s: String) -> [i32; 28] {
            let mut res = [0; 28];

            for ch in s.bytes() {
                res[(ch - b'a') as usize] += 1;
            }

            res
        }

        let s_count = count_chars(s);
        let t_count = count_chars(t);

        let mut res = 0;

        for i in 0..s_count.len() {
            let diff = s_count[i] - t_count[i];

            if diff > 0 {
                res += diff;
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
    fn min_steps() {
        assert_eq!(Solution::min_steps("bab".to_owned(), "aba".to_owned()), 1);
        assert_eq!(
            Solution::min_steps("anagram".to_owned(), "mangaar".to_owned()),
            0
        );
        assert_eq!(
            Solution::min_steps("leetcodez".to_owned(), "zpractice".to_owned()),
            5
        );
    }
}
