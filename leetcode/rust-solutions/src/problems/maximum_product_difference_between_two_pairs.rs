struct Solution;

impl Solution {
    pub fn max_product_difference(nums: Vec<i32>) -> i32 {
        let mut a = vec![0, 0, i32::MAX, i32::MAX];

        for n in nums {
            if n >= a[0] {
                a[1] = a[0];
                a[0] = n;
            } else {
                if n > a[1] {
                    a[1] = n;
                }
            }

            if n <= a[2] {
                a[3] = a[2];
                a[2] = n;
            } else {
                if n < a[3] {
                    a[3] = n;
                }
            }
        }

        a[0] * a[1] - a[2] * a[3]
    }
}
