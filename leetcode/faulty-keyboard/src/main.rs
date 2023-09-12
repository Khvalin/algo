struct Solution;

impl Solution {
    pub fn final_string(s: String) -> String {
        let mut prefix = vec![];
        let mut postfix = vec![];

        let mut c = 0;

        for ch in s.chars().rev() {
            if ch == 'i' {
                c += 1;
                continue;
            }

            if c % 2 == 1 {
                prefix.push(ch);
            } else {
                postfix.push(ch)
            }
        }

        postfix.reverse();

        let res = format!("{}{}", prefix.into_iter().collect::<String>(), postfix.into_iter().collect::<String>());
        res
    }
}

fn main() {
    let res = Solution::final_string("abicdiefi".to_owned());
    assert_eq!("febacd", res);
    println!("All passed!");
}
