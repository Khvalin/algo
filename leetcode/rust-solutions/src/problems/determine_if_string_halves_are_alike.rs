struct Solution;

use std::collections::HashSet;

impl Solution {
    pub fn halves_are_alike(s: String) -> bool {
        assert!(s.len() % 2 == 0);

        let mut sum = 0;
        let vowels = HashSet::from(['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U']);

        for (i, ch) in (0..s.len()).zip(s.chars()) {
            let mut d = 0;
            if vowels.contains(&ch) {
                if i < s.len() >> 1 {
                    d = 1;
                } else {
                    d = -1;
                }
            }

            sum += d;
        }

        sum == 0
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_halves_are_alike_empty_string() {
        assert!(Solution::halves_are_alike(String::from("")));
    }

    #[test]
    fn test_halves_are_alike_equal_halves() {
        assert!(Solution::halves_are_alike(String::from("abcABC")));
        assert!(Solution::halves_are_alike(String::from("aAaA")));
    }

    #[test]
    fn test_halves_are_alike_not_equal_halves() {
        assert!(!Solution::halves_are_alike(String::from("aabcdABC")));
        assert!(!Solution::halves_are_alike(String::from("aeiouAEIOZ")));
    }

    #[test]
    #[should_panic]
    fn test_halves_are_alike_odd_length_string() {
        // This test should panic due to the assertion in the function
        Solution::halves_are_alike(String::from("abc"));
    }
}
