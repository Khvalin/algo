struct Solution;
use std::collections::HashMap;
use std::usize;

#[derive(Debug)]
struct Node {
    chars: HashMap<char, Node>,
    val: Option<usize>,
}

#[derive(Debug)]
struct Trie {
    root: Node,
}

impl Trie {
    fn new() -> Trie {
        Trie {
            root: Node {
                chars: HashMap::new(),
                val: None,
            },
        }
    }
    fn add_string(&mut self, string: String, val: usize) {
        let mut current_node = &mut self.root;
        for c in string.chars() {
            current_node = moving(current_node).chars.entry(c).or_insert(Node {
                chars: HashMap::new(),
                val: None,
            });
        }
        current_node.val = Some(val);
    }
}

fn moving<T>(t: T) -> T {
    t
}

#[macro_export]
macro_rules! svec {
    () => {{
        let v = Vec::<String>::new();
        v

    }};
    ($($elem:expr),+ $(,)?) => {{
        let v = vec![
            $( String::from($elem), )*
        ];
        v
    }};
}
impl Solution {
    fn suggested_products(mut products: Vec<String>, search_word: String) -> Vec<Vec<String>> {
        let mut res = vec![];
        let mut trie = Trie::new();

        for (i, s) in products.iter().enumerate() {
            trie.add_string(s.to_owned(), i)
        }

        let mut nodes = vec![&trie.root];

        for ch in search_word.chars() {
            let mut next = vec![];
            for _ in 0..nodes.len() {
                let node = &nodes.pop().unwrap();
                if node.chars.contains_key(&ch) {
                    next.push(&node.chars[&ch]);
                }
            }
            nodes = next;

            let mut candidates = nodes.clone();

            let mut list = vec![];
            let mut j = 0;

            while list.len() < 3 && j < candidates.len() {
                for c in 'a'..'z' {
                    if candidates[j].chars.contains_key(&c) {
                        if j < candidates.len()-1 {
                            candidates[j+1] = candidates[j];
                        }
                        candidates[j] = &candidates[j].chars[&c]
                    }

                    if candidates[j].val.is_some() {
                        let ind = candidates[j].val.unwrap();
                        list.push(products[ind].clone());
                    }
                }

                if candidates[j].chars.len() == 0 {
                    j += 1;
                }
            }

            res.push(list);
        }

        res
    }
}

fn main() {
    println!("Hello, world!");
}

#[test]
fn test() {
    let products = svec!["mobile", "mouse", "moneypot", "monitor", "mousepad"];
    let search_word = "mouse1".to_string();
    let res: Vec<Vec<String>> = vec![
        svec!["mobile", "moneypot", "monitor"],
        svec!["mobile", "moneypot", "monitor"],
        svec!["mouse", "mousepad"],
        svec!["mouse", "mousepad"],
        svec!["mouse", "mousepad"],
    ];
    assert_eq!(Solution::suggested_products(products, search_word), res);
    let products = svec!["havana"];
    let search_word = "havana".to_string();
    let res: Vec<Vec<String>> = vec![
        svec!["havana"],
        svec!["havana"],
        svec!["havana"],
        svec!["havana"],
        svec!["havana"],
        svec!["havana"],
    ];
    assert_eq!(Solution::suggested_products(products, search_word), res);
    let products = svec!["bags", "baggage", "banner", "box", "cloths"];
    let search_word = "bags".to_string();
    let res: Vec<Vec<String>> = vec![
        svec!["baggage", "bags", "banner"],
        svec!["baggage", "bags", "banner"],
        svec!["baggage", "bags"],
        svec!["bags"],
    ];
    assert_eq!(Solution::suggested_products(products, search_word), res);
    let products = svec!["havana"];
    let search_word = "tatiana".to_string();
    let res: Vec<Vec<String>> = vec![
        svec![],
        svec![],
        svec![],
        svec![],
        svec![],
        svec![],
        svec![],
    ];
    assert_eq!(Solution::suggested_products(products, search_word), res);
}
