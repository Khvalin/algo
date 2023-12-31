struct Solution;

impl Solution {
    pub fn max_length_between_equal_characters(s: String) -> i32 {
        let mut res: i32 = -1;
        let mut last = [-1; 30];
        let mut bytes = s.bytes();

        for i in 0..s.len() {
            let ch = bytes.next().unwrap();
            let ind = (ch - b'a') as usize;
            if last[ind] < 0 {
                last[ind] = i as i32;
                continue;
            }
            res = res.max(i as i32 - last[ind] - 1);
        }
        res
    }
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn max_length_between_equal_characters() {
        assert_eq!(
            Solution::max_length_between_equal_characters("aa".to_owned()),
            0
        );
    }
}
