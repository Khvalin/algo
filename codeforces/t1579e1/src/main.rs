use std::collections::VecDeque;
use std::io;
use std::io::prelude::*;
use std::io::BufWriter;
use std::io::{stdout, Write};

fn read_input() -> Vec<Vec<i32>> {
    let stdin = io::stdin();
    // Read one line of input iterator-style
    let mut iter = stdin.lock().lines();
    let input = iter.next().unwrap().unwrap();
    let t = input.parse::<i32>().unwrap();

    let mut res: Vec<Vec<i32>> = vec![];
    for _ in 0..t {
        iter.next();
        let s = iter.next().unwrap().unwrap();
        let numbers: Vec<i32> = s
            .split_whitespace()
            .map(|d| d.parse::<i32>().unwrap())
            .collect();
        res.push(numbers);
    }

    return res;
}

fn solve(t: Vec<i32>) -> Vec<i32> {
    let mut dq: VecDeque<i32> = VecDeque::new();

    for n in t {
        if dq.len() == 0 || n > *dq.front().unwrap() {
            dq.push_back(n);
            continue;
        }

        dq.push_front(n);
    }

    let mut res = vec![0; dq.len()];

    for (i, n) in dq.iter().enumerate() {
        res[i] = *n
    }

    res
}

fn main() {
    let mut buffer = BufWriter::new(stdout());

    let a = read_input();
    for e in a {
        let nums = solve(e);

        let s = nums
            .iter()
            .map(|n| n.to_string())
            .collect::<Vec<String>>()
            .join(" ");
        write!(buffer, "{}\n", s);
    }

    buffer.flush().unwrap();
}
