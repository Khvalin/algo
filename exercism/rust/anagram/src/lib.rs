use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &'a str, possible_anagrams: &'a [&str]) -> HashSet<&'a str> {
    let mut result: HashSet<&str> = HashSet::new();
    let w1 = word.to_lowercase();

    let mut needle = w1.chars().collect::<Vec<char>>();
    needle.sort_unstable();

    'outer: for s in possible_anagrams {
        let w2 = s.to_lowercase();
        if w1 == w2 {
            continue;
        }
        if w1.len() != w2.len() {
            continue;
        }

        let mut b = w2.chars().collect::<Vec<char>>();
        b.sort_unstable();

        for (i, ch) in b.iter().enumerate() {
            if *ch != needle[i] {
                continue 'outer;
            }
        }

        result.insert(s);
    }

    result
}
