
pub fn abbreviate(phrase: &str) -> String {
    let mut res = vec![];
    let mut prev = ' ';
    for ch in phrase.chars() {
        if ch == '\'' {
            continue;
        }
        if ch.is_alphabetic() && (!prev.is_alphabetic() || prev.is_lowercase() && ch.is_uppercase())
        {
            res.push(ch.to_uppercase().to_string());
        }
        prev = ch;
    }

    res.join("")
}
