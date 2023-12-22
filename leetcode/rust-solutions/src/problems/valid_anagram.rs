struct Solution;

impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        if s.len() != t.len() {
            return false;
        }

        let a_code = 'a' as usize;
        let mut counts = [0; 30];

        let mut iter = t.bytes().into_iter();

        for ch in s.bytes() {
            counts[ch as usize - a_code] += 1;

            let tch = iter.next().unwrap();
            counts[tch as usize - a_code] -= 1;
        }

        counts.into_iter().all(|c| *c == 0)
    }
}
