struct Solution;


impl Solution {
    pub fn sum_even_after_queries(mut nums: Vec<i32>, queries: Vec<Vec<i32>>) -> Vec<i32> {
        let mut sum = (&nums).into_iter().fold(0, |acc, n| {
            if n % 2 == 0 {
                return acc + n;
            }
            acc
        });

        let mut res = vec![];
        for q in queries {
            let ind = q[1] as usize;
            if nums[ind] % 2 == 0 {
                sum -= nums[ind]
            }

            nums[ind] += q[0];

            if nums[ind] % 2 == 0 {
                sum += nums[ind]
            }

            res.push(sum);
        }

        res
    }
}



fn main() {
    println!("Hello, world!");
}
