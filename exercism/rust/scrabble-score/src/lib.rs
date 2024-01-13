use unicode_segmentation::Graphemes;
use std::collections::HashMap;

pub fn score(input: &'static str) -> usize {
  let mut scores: HashMap<char, usize> = HashMap::new();

  insert_scores(&mut scores, vec!['Q', 'Z'], 10);
  insert_scores(&mut scores, vec!['J', 'X'], 8);
  insert_scores(&mut scores, vec!['K'], 5);
  insert_scores(&mut scores, vec!['F', 'H', 'V', 'W', 'Y'], 4);
  insert_scores(&mut scores, vec!['B', 'C', 'M', 'P'], 3);
  insert_scores(&mut scores, vec!['D', 'G'], 2);
  insert_scores(&mut scores, vec!['A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'], 1);

  UnicodeSegmentation::graphemes(input.to_string().to_uppercase())
    .fold(0, |sum, ch| sum + scores.get(&ch).unwrap_or(&0))
  
}


fn insert_scores(scores: &mut HashMap<char, usize>, letters: Vec<char>, score: usize) -> () {
  for ch in letters {
    scores.insert(ch, score);
  }
}
