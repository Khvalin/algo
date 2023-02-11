use std::io::{self, prelude::*, BufReader};

fn main() {
    let pi = "314159265358979323846264338327";
    let buffer = BufReader::new(io::stdin());
    let mut input_iter = buffer.lines();
    input_iter.next();

    for line in input_iter {
        let mut len = 0;
        for ch in line.unwrap().chars() {
            if len >= pi.len() || ch != pi.chars().nth(len).unwrap() {
                break;
            }
            len += 1;
        }
        println!("{}", len);
    }
}
