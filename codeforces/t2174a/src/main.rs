use std::{
    collections::HashMap,
    io::{self, Write},
};

fn solve(s: Vec<char>, t: Vec<char>) -> String {
    let impossible = "Impossible".to_owned();
    let mut result: Vec<char> = Vec::new();

    let mut tcount: HashMap<char, u32> = HashMap::new();

    for ch in &t {
        tcount.entry(*ch).and_modify(|e| *e += 1).or_insert(1);
    }

    for ch in &s {
        if (!tcount.contains_key(ch)) || (*tcount.get(ch).unwrap() == 0) {
            return impossible;
        }

        tcount.entry(*ch).and_modify(|e| *e -= 1);
    }

    let mut i = 0;
    for _ in 0..t.len() {
        let mut c = '!';
        for ch in 'a'..='z' {
            if i < s.len() && ch == s[i] {
                c = ch;
                i += 1;
                break;
            }

            if *tcount.get(&ch).unwrap_or(&0) > 0 {
                c = ch;
                tcount.entry(ch).and_modify(|e| *e -= 1);
                break;
            }
        }

        if c == '!' {
            return impossible;
        }
        result.push(c);
    }

    result.into_iter().collect()
}

fn main() {
    let mut line = String::new();
    io::stdin().read_line(&mut line).unwrap();

    let total = line.trim().parse::<i32>().unwrap();

    for _ in 0..total {
        line.clear();
        io::stdin().read_line(&mut line).unwrap();
        let s: Vec<char> = line.trim().chars().collect();

        line.clear();
        io::stdin().read_line(&mut line).unwrap();
        let t: Vec<char> = line.trim().chars().collect();

        io::stdout()
            .write_fmt(format_args!("{}\n", solve(s, t)))
            .unwrap();
    }
}
