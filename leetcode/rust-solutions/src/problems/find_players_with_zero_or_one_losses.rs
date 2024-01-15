use std::collections::HashMap;

struct Solution;

impl Solution {
    /// https://leetcode.com/problems/find-players-with-zero-or-one-losses/
    pub fn find_winners(matches: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut losses: HashMap<i32, i32> = HashMap::new();

        for m in matches {
            losses.insert(m[0], *losses.get(&m[0]).unwrap_or(&0));
            losses.insert(m[1], 1 + *losses.get(&m[1]).unwrap_or(&0));
        }

        let mut res = vec![vec![], vec![]];

        for (k, v) in losses {
            if v < 2 {
                res[v as usize].push(k);
            }
        }

        res[0].sort_unstable();
        res[1].sort_unstable();

        res
    }
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn find_winners() {
        assert_eq!(
            Solution::find_winners(vec![vec![1, 2]]),
            vec![vec![1], vec![2]]
        );
    }
}
