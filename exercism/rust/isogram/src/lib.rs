use std::collections::HashSet;
use unicode_segmentation::UnicodeSegmentation;

pub fn check(candidate: &str) -> bool {
    let mut seen = HashSet::new();
    for grapheme in candidate.to_lowercase().graphemes(true) {
        if grapheme == " " || grapheme == "-" {
            continue;
        }

        if seen.contains(grapheme) {
            return false;
        }

        seen.insert(grapheme);
    }

    true

}
