use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn count_characters(words: Vec<String>, chars: String) -> i32 {
        let mut m: HashMap<char, usize> = HashMap::new();

        for ch in chars.chars() {
            *m.entry(ch).or_insert(0) += 1;
        }

        let res = words.into_iter().fold(0, |s, w| {
            let mut wm: HashMap<char, usize> = HashMap::new();

            for ch in w.chars() {
                *wm.entry(ch).or_insert(0) += 1;
            }

            let mut f = true;

            for (key, value) in wm {
                f = f && (value <= *m.entry(key).or_default());
            }

            if f {
                return s + w.len() as i32;
            }

            s
        });

        res
    }
}

fn main() {
    println!("Hello, world!");
}
