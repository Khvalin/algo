struct Solution;

impl Solution {
    pub fn max_score(s: String) -> i32 {
        let mut ones = 0;
        let mut chars = s.chars();
        for ch in s.chars() {
            if ch == '1' {
                ones += 1;
            }
        }

        let mut chars = s.chars();
        let mut zeroes = 0;
        let mut res = i32::MIN;
        for _ in 0..(s.len() - 1) {
            let ch = chars.next().unwrap();
            if ch == '0' {
                zeroes += 1;
            } else {
                ones -= 1;
            }

            res = res.max(zeroes + ones)
        }
        res
    }
}
