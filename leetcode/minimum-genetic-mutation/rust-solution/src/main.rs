use std::collections::HashSet;

struct Solution;

impl Solution {
    pub fn min_mutation(start: String, end: String, bank: Vec<String>) -> i32 {
        let mut res = bank.len() as i32 + 10;

        let mut q = vec![(&start, HashSet::new() as HashSet<usize>, 0)];

        while q.len() > 0 {
            let state = q.pop().unwrap();

            let cur = state.0;
            let visited = &(state.1);
            let dist = state.2;

            if *cur == end {
                res = res.min(dist);
                continue;
            }

            for i in 0..bank.len() {
                let dna = &bank[i];
                let mut d = 0;
                for j in 0..dna.len() {
                    if dna.as_bytes()[j] != cur.as_bytes()[j] {
                        d += 1
                    }
                }

                if (*visited).contains(&i) || d != 1 {
                    continue;
                }

                let mut copy = HashSet::new();
                copy.extend(visited);
                copy.insert(i);

                q.push((dna, copy, dist + 1));
            }
        }

        if res > bank.len() as i32 {
            return -1;
        }

        return res;
    }
}

fn main() {
    let res = Solution::min_mutation(
        "AACCGGTT".to_string(),
        "AACCGGTA".to_string(),
        vec!["AACCGGTA".to_string()],
    );
    println!("{:?}", res);
}
