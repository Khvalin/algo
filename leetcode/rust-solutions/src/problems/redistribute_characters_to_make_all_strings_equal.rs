struct Solution;

impl Solution {
    pub fn make_equal(words: Vec<String>) -> bool {
        let mut counts = [0; 30];
        for w in &words {
            for ch in w.bytes() {
                counts[(ch - b'a') as usize] += 1;
            }
        }

        counts.into_iter().all(|c| c % &words.len() == 0)
    }
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn make_equal() {
        let words: Vec<String> = vec![
            String::from("abc"),
            String::from("aabc"),
            String::from("bc"),
        ];
        assert_eq!(Solution::make_equal(words), true);
    }
}
