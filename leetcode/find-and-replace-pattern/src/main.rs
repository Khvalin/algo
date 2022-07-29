struct Solution;

impl Solution {
    pub fn find_and_replace_pattern(words: Vec<String>, p: String) -> Vec<String> {
        let mut res = vec![];

        let mut a = vec![None; ('z' as usize)+1];
        let mut b = vec![None; ('z' as usize)+1];

        for w in words {
            a.fill(None);
            b.fill(None);

            let mut f = true;

            let mut piter = p.chars();


            for ch in w.chars() {
                let pch = piter.next().unwrap();
                let j = ch as usize;
                let k = pch as usize;
                if a[j].is_none() && b[k].is_none() {
                    a[j] = Some(pch);
                    b[k] = Some(ch);
                    continue;
                }

                if a[j] == Some(pch) && b[k] == Some(ch) {
                    continue;
                }

                f = false;
                break;
            }

            if f {
                res.push(w);
            }
        }

        res
    }
}
