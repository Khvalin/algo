struct Solution;

impl Solution {
    pub fn final_string(s: String) -> String {
        let mut res = s.chars().map(|_| '\0').collect::<Vec<char>();
        let mut c = 0;
        let l = 0;
        let r = s.len() - 1;

        for ch in s.chars().rev() {
            if ch == 'i' {
                c += 1;
                continue;
            }

            if c % 2 == 1 {
                res[l] = ch;
                l += 1;
            } else {
                res[r] = ch;
                r -= 1;
            }
        }

        res
    }
}

fn main() {
    let res = Solution::final_string("abicdiefi".to_owned());
    assert_eq!("febacd", res);
    println!("All passed!");
}
