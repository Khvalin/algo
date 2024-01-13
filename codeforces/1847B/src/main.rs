use std::io::{self, BufRead};


fn main() {
    
    let stdin = io::stdin();
    let mut iterator = stdin.lock().lines();
    let n = iterator.next().unwrap().unwrap().parse::<i64>();
    for _i in 0..n.unwrap() {
        iterator.next();
        let line = iterator.next().unwrap().unwrap();

        let nums = line
            .split_ascii_whitespace()
            .filter_map(|s| s.parse::<i64>().ok())
                .collect::<Vec<_>>();

        let mut m = nums[0];
        for n in &nums {
            m &= n;
        }

        if m > 0 {
            println!("1");
            continue;
        }

        let mut c = 1;
        let mut m: Option<i64> = None;
        for n in &nums {
            if (&m).is_none() {
                m = Some(*n);
                continue;
            }

            let mut t = m.unwrap();
            t &= n;
            if t == 0 {
                c+=1;
                m = None;
                continue;
            }
            
            m = Some(t);
        }
        
        if m.is_some() {
            c -= 1;
        }

        println!("{}", c.max(1))
    }
}
