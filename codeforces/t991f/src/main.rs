
fn read_input() -> String {
    let mut buffer = String::new();
    std::io::stdin().read_line(&mut buffer).expect("Failed");
    buffer
}

fn solve(n:u128, d:u32) -> String {
    let mut res = n.to_string();
    if n < 100 || d > 1 {
        return res;
    }

    let f = n as f64;


    for p in 2..15 {
        let base = f.powf(1.0/(p as f64)) as u128;
        let pow = base.pow(p);

        if pow == 0 {
            continue;
        }

        let mul = n / pow;
        let diff = n - mul * pow;

        let mut s = "".to_owned();


        if base > 1 {
            s = base.to_string();
            if p > 1 {
                s = format!("{}^{}", s, p);
            }
        }

        if mul > 1 {
            if s.len() > 0 {
                s = format!("*{}", s);
            }
            s = format!("{}{}", mul, s);
        }
        if diff > 0 {
            if s.len() > 0 {
                s = format!("{}+", s);
            }
            s = format!("{}{}", s, solve(diff, d+1));
        }

        dbg!(s.to_string());
        if s.len() <= res.len() {
            res = s.to_owned();
        }
    }

    res
}

fn main() {
    let n = read_input().trim().parse::<u128>().unwrap();

    let res = solve(n, 0);


    println!("{}", res);
}
