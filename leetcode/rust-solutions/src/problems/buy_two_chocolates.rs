struct Solution;

impl Solution {
    pub fn buy_choco(prices: Vec<i32>, money: i32) -> i32 {
        let mut a = [i32::MAX, i32::MAX];

        for p in prices {
            if p <= a[0] {
                a[1] = a[0];
                a[0] = p;
                continue;
            }

            if p < a[1] {
                a[1] = p;
            }
        }

        if a[0] + a[1] > money {
            return money;
        }

        money - (a[0] + a[1])
    }
}
