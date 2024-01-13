use std::collections::HashMap;

pub fn word_count(input: &str) -> HashMap<String, u32> {
    let mut counts: HashMap<String, u32> = HashMap::new();

    let s = input.to_lowercase();

    let mut w = vec![];
    for i in 0..=s.len() {
        let mut ch = '\0'; // treat end of string as a whitespace
        if i < s.len() {
            ch = s.as_bytes()[i] as char;
        }

        if ch.is_alphanumeric() || ch == '\'' {
            // ch is not a whitespace
            w.push(ch);
            continue;
        }

        if w.is_empty() {
            continue;
        }

        // we have a match, so count it
        let mut word = String::from_iter(&w);
        word = word.trim_matches('\'').to_string();

        if word.is_empty() {
            continue;
        }
        let c = 1 + counts.get(&word).unwrap_or(&0);
        counts.insert(word, c);

        w.clear();
    }

    counts
}
