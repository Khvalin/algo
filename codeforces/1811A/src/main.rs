use std::io::{self, prelude::*, stdout, BufReader};

fn main() {
    let buffer = BufReader::new(io::stdin());
    let mut input_iter = buffer.lines();

    input_iter.next();

    let mut d: char = '\0';
    for line in input_iter {
        let numbers: Vec<_> = line.unwrap().split_whitespace().map(String::from).collect();

        if numbers.len() == 2 {
            d = numbers.last().unwrap().as_str().chars().last().unwrap();
            continue;
        }

        let num = numbers.first().unwrap();

        for ch in num.chars() {
            if ch < d {
                print!("{}", d);
                d = '\0'
            }
            print!("{}", ch);
        }
        if d != '\0' {
            print!("{}", d);
        }

        println!();
    }

    stdout().flush().unwrap_or_default();
}
