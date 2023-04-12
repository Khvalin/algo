use std::io::{self, prelude::*, BufReader};

fn solve(nums: Vec<i32> ) -> i64 {
    let mut sum = 0 as i64;
    let mut c = 0;
    let mut min = nums[0].abs();

    for n in nums.iter() {
        sum += n.abs() as i64;
        min = min.min(n.abs());
        if *n < 0 {
            c += 1;
        }
    }

    if c % 2 == 1 {
        sum -= 2*min as i64;
    }

    return sum as i64;
}

fn main() {
    let buffer = BufReader::new(io::stdin());
    let mut input_iter = buffer.lines();

    input_iter.next();

    for line in input_iter {
        let numbers: Vec<i32> = line.unwrap()
        .split_whitespace()
        .map(|s| s.parse().expect("parse error"))
        .collect();

        if numbers.len() == 1 {
            continue;
        }
        let res = solve(numbers);
        println!("{}", res);
    }
}
