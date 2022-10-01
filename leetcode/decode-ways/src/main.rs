struct Solution;

impl Solution {
    pub fn num_decodings(s: String) -> i32 {
        let mut state = [0, 0, 0];
        let mut prev = '-'; // to keep track of previously visited char

        let mut i = s.len();

        for ch in s.chars().rev() {
            i -= 1;
            let mut res = 0;

            // one-digit numbers
            if ch != '0' {
                if i == s.len() - 1 {
                    res += 1;
                } else {
                    res += state[(i + 1) % 3];
                }
            }

            // two-digit numbers
            if ch == '1' || (ch == '2' && prev < '7') {
                if i == s.len() - 2 {
                    res += 1
                } else {
                    res += state[(i + 2) % 3];
                }
            }

            state[i % 3] = res;

            prev = ch;
        }

        state[0]
    }
}

fn main() {
    println!("{}", Solution::num_decodings("12".to_string()));
    println!("{}", Solution::num_decodings("12121212121212".to_string()));
}
