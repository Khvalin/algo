struct Solution;

impl Solution {
    pub fn min_operations(s: String) -> i32 {
        let mut next = [0, 1];
        let mut res = [0, 0];

        for ch in s.chars() {
            let mut n = 0;
            if ch == '1' {
                n = 1;
            }

            for i in 0..2 {
                if n != next[i] {
                    res[i] += 1;
                }

            }

            next[0] ^= 1;
            next[1] ^= 1;
        }

        res[0].min(res[1])
    }
}
