use std::io::{self, BufRead, BufReader};

fn solve(chars: &mut Vec<char>) {
    let mut count: [u8; 26] = [0; 26];

    for ch in chars.iter() {
        let ord = *ch as u8 - 'a' as u8;
        count[ord as usize] += 1;
    }

    let mut i = 0;
    for p in 0..26 {
        while count[p] > 0 {
            chars[i] = ('a' as u8 + p as u8) as char;
            i+=1;
            count[p] -= 1;
        }
    }
}

fn main() {
    let stdin = io::stdin();
    let reader = BufReader::new(stdin);

    let mut first_line = true;
    for line in reader.lines() {
        match line {
            Ok(content) => {
                if first_line {
                    first_line = false;
                    continue;
                }
                let mut chars:Vec<char> = content.chars().collect();
                solve(&mut chars);
                for &c in &chars {
                    print!("{}", c);
                }
                println!();
            }
            Err(_) => return,
        }
    }
}
