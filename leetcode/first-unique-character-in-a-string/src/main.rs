struct Solution;

impl Solution {
    pub fn first_uniq_char(s: String) -> i32 {
        let mut res = s.len() + 1;
        for p in 'a'..='z' {
            let mut count = 0;
            let mut ind = s.len() + 1;
            for (i, ch) in s.chars().enumerate() {
                if ch == p {
                    count += 1;
                    ind = i;
                }
            }

            if count == 1 {
                res = res.min(ind);
            }
        }

        if res > s.len() {
            return -1;
        }

        res as i32
    }
}

fn main() {
    println!("{}", Solution::first_uniq_char("eleetcode".to_owned()));
}
