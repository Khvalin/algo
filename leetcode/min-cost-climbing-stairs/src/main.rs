struct Solution {}

impl Solution {
    pub fn min_cost_climbing_stairs(cost: Vec<i32>) -> i32 {
        let mut a = 0;
        let mut b = 0;

        for mut n in cost {
            n += a.min(b);
            a = b;
            b = n;
        }

        a.min(b)
    }
}

fn main() {
    println!("Hello, world!");
}

#[test]
fn test() {
    let cost: Vec<i32> = vec![1, 2, 1];
    assert_eq!(Solution::min_cost_climbing_stairs(cost), 2);

    assert_eq!(
        Solution::min_cost_climbing_stairs(vec![1, 100, 1, 1, 1, 100, 1, 1, 100, 1]),
        6
    );
}
