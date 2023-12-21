struct Solution;

impl Solution {
    pub fn max_width_of_vertical_area(points: Vec<Vec<i32>>) -> i32 {
        let mut a: Vec<i32> = points.iter().map(|p| p[0]).collect();
        a.sort_unstable();

        let mut res = -1;

        for i in 1..a.len() {
            res = res.max(a[i] - a[i - 1]);
        }

        res
    }
}
