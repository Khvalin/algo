use std::io::{self, BufRead};

// fn parse_int(s: String) -> i32 {
//     s.bytes()
//         .fold(0, |acc, curr| 10 * acc + (curr - b'0') as i32)
// }

fn main() {
    let stdin = io::stdin();
    let mut iterator = stdin.lock().lines();
    let n = iterator.next().unwrap().unwrap().parse::<i64>();
    for _i in 0..n.unwrap() {
        iterator.next();
        let line = iterator.next().unwrap().unwrap();

        let numbers = line
            .split_ascii_whitespace()
            .filter_map(|s| s.parse::<i64>().ok())
            .collect::<Vec<_>>();

        let mut f = true;
        let mut p = 0;
        for n in numbers {
            if n > 0 {
                if p > 0 {
                    f = false;
                    break;
                }
                p = 0;
            }
            p += n;
        }

        println!("{}", if f { "YES" } else { "NO" })
    }
}
