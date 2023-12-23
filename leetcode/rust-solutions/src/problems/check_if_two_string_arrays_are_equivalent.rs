struct Solution;

impl Solution {
    pub fn array_strings_are_equal(word1: Vec<String>, word2: Vec<String>) -> bool {
        let mut i = 0;
        let mut j = 0;

        let mut char_iter1 = "".chars();
        let mut char_iter2 = "".chars();

        let mut ch: (Option<char>, Option<char>) = (None, None);

        while i <= word1.len() && j <= word2.len() {
            if ch.0.is_none() {
                ch.0 = char_iter1.next();
            }

            if ch.1.is_none() {
                ch.1 = char_iter2.next();
            }

            if ch.0.is_none() {
                if i < word1.len() {
                    char_iter1 = word1[i].chars();
                }
                i += 1;
            }

            if ch.1.is_none() {
                if j < word2.len() {
                    char_iter2 = word2[j].chars();
                }
                j += 1;
            }

            if ch.0.is_none() || ch.1.is_none() {
                continue;
            }

            if ch.0.unwrap() != ch.1.unwrap() {
                return false;
            }

            ch = (None, None)
        }

        (i == 1 + word1.len()) && (j == 1 + word2.len())
    }
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn test_array_strings_are_equal() {
        assert_eq!(
            Solution::array_strings_are_equal(
                vec!["ab".to_owned(), "c".to_owned()],
                vec!["a".to_owned(), "bc".to_owned()]
            ),
            true
        );
    }
}
