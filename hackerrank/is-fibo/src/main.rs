use std::env;
use std::fs::File;
use std::io::{self, BufRead, Write};

/*
 * Complete the 'isFibo' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts LONG_INTEGER n as parameter.
 */

fn isFibo(n: i64) -> String {
    let mut a = 0;
    let mut b = 1;

    while a < n {
        let c = a + b;
        a = b;
        b = c;
    }

    if a == n {
        return String::from("IsFibo");
    }

    return String::from("IsNotFibo");
}

fn main() {
    let stdin = io::stdin();
    let mut stdin_iterator = stdin.lock().lines();

    let mut fptr = File::create(env::var("OUTPUT_PATH").unwrap()).unwrap();

    let t = stdin_iterator.next().unwrap().unwrap().trim().parse::<i32>().unwrap();

    for _ in 0..t {
        let n = stdin_iterator.next().unwrap().unwrap().trim().parse::<i64>().unwrap();

        let result = isFibo(n);

        writeln!(&mut fptr, "{}", result).ok();
    }
}
