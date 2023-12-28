struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn solve(
        memo: &mut HashMap<(usize, u8, usize, i32), i32>,
        s: &Vec<u8>,
        cur: usize,
        last_char: u8,
        last_char_count: usize,
        k: i32,
    ) -> i32 {
        if k < 0 {
            return i32::MAX;
        }
        if cur >= s.len() {
            return 0;
        }

        let key = (cur, last_char, last_char_count, k);
        if memo.contains_key(&key) {
            return *memo.get(&key).unwrap();
        }

        let mut res = i32::MAX;
        if s[cur] == last_char {
            let mut inc = 0;
            if last_char_count == 1 || last_char_count == 9 || last_char_count == 99 {
                inc = 1;
            }
            res = inc + Solution::solve(memo, s, cur + 1, last_char, last_char_count + 1, k);
        } else {
            //keep char
            res = res.min(1 + Solution::solve(memo, s, cur + 1, s[cur], 1, k));

            if k > 0 {
                //delete char
                res = res.min(Solution::solve(
                    memo,
                    s,
                    cur + 1,
                    last_char,
                    last_char_count,
                    k - 1,
                ));
            }
        }

        memo.insert(key, res);

        res
    }

    pub fn get_length_of_optimal_compression(s: String, k: i32) -> i32 {
        let mut memo = HashMap::new();
        Solution::solve(&mut memo, &s.into_bytes(), 0, b'\0', 0, k)
    }
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn get_length_of_optimal_compression() {
        assert_eq!(
            Solution::get_length_of_optimal_compression("aaabcccd".to_owned(), 2),
            4
        );
        assert_eq!(
            Solution::get_length_of_optimal_compression("aabbaa".to_owned(), 2),
            2
        );
        assert_eq!(
            Solution::get_length_of_optimal_compression("jifgweftpsnlorhrlsgdpddlcfabislqhgkkedtgolfpowdcknsjflrjlnqoiriuefxlsvufjqpvuimehuqheakuufbnahdplopo".to_owned(), 34),
            60
        );
    }
}
