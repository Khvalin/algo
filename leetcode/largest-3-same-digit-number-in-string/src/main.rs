struct Solution;

impl Solution {
    pub fn largest_good_integer(num: String) -> String {
        let mut res = '\0';
        let chars = num.chars().collect::<Vec<char>>();

        for i in 1..chars.len() - 1 {
            let ch = chars[i];
            if ch == chars[i - 1] && ch == chars[i + 1] && ch > res {
                res = ch;
            }
        }

        if res < '0' {
            return "".to_owned();
        }

        format!("{0}{0}{0}", res)
    }
}

fn main() {
    println!("{}", Solution::largest_good_integer("655546".to_owned()));
    println!("{}", Solution::largest_good_integer("000".to_owned()));
    println!("{}", Solution::largest_good_integer("29995551".to_owned()));
}
