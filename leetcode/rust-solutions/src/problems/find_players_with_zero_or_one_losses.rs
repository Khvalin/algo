use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn find_winners(matches: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut loses:HashMap<i32, i32> = HashMap::new();

        let mut min = 100_001;
        let mut max = 0;

        for m in matches {
            loses.insert(m[0], *loses.get(&m[0]).unwrap_or(&0));
            loses.insert(m[1], 1 + *loses.get(&m[1]).unwrap_or(&0));

            min = min.min(m[0]).min(m[1]);
            max = max.max(m[0]).max(m[1]);
        }

        let mut res = vec![vec![], vec![]];

        for (k, v) in loses.keys().zip(loses.values()) {
            if *v < 2 {
                res[*v as usize].push(*k);
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
