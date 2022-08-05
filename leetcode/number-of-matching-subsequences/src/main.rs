struct Solution;

impl Solution {
    pub fn num_matching_subseq(s: String, words: Vec<String>) -> i32 {
        let mut next: Vec<Vec<usize>> = vec![];
        let mut counter: Vec<usize> = words.to_owned().into_iter().map(|_| 0).collect();

        for _ in 0..=('z' as usize) {
            let a: Vec<usize> = vec![];
            next.push(a);
        }

        for i in 0..words.len() {
            let ch = words[i].chars().nth(0).unwrap();

            next[ch as usize].push(i);
        }
        let mut res = 0;

        s.chars().into_iter().for_each(|ch| {
            let a = next[ch as usize].to_owned();
            next[ch as usize].clear();

            for i in a {
                counter[i] += 1;
                if counter[i] >= words[i].len() {
                    res += 1;
                    continue;
                }

                let c = words[i].chars().nth(counter[i]).unwrap();

                next[c as usize].push(i);
            }
        });

        res
    }
}

fn main() {
    let res = Solution::num_matching_subseq(
        "abcdeb".to_string(),
        vec![
            "a".to_string(),
            "bb".to_string(),
            "acd".to_string(),
            "ace".to_string(),
            "ahce".to_string(),
        ],
    );

    println!("{}", res);
}
