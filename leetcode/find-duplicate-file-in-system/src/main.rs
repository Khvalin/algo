use std::collections::HashMap;


struct Solution;

impl Solution {
    pub fn find_duplicate(paths: Vec<String>) -> Vec<Vec<String>> {
        let mut m: HashMap<String, Vec<String>> = HashMap::new();

        for p in &paths {
            let parts: Vec<&str> = p.split(' ').collect();
            let folder = parts[0];

            for i in 1..parts.len() {
                let fragments:Vec<&str> = parts[i].split('(').collect();
                let filename = format!("{}/{}", folder, fragments[0]);
                m.entry(fragments[1].to_string()).or_insert(vec!()).push(filename);
            }
        }

        let mut res = vec![];
        for a in m.values() {
            if a.len() > 1 {
                res.push(a.to_vec());
            }
        }
        return res;
    }
}

fn main() {
    println!("Hello, world!");
}
