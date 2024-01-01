struct Solution;

impl Solution {
    pub fn find_content_children(mut g: Vec<i32>, mut s: Vec<i32>) -> i32 {
        let mut res = 0;
        g.sort_unstable();
        s.sort_unstable();

        let mut j = 0;

        for n in &g {
            while j < s.len() && s[j] < *n {
                j += 1;
            }

            if j < s.len() && s[j] >= *n {
                res += 1;
            }
            j += 1;
        }

        res
    }
}
