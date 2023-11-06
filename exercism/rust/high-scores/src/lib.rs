#[derive(Debug)]
pub struct HighScores<'a> {
    scores: &'a [u32],
}

impl<'a> HighScores<'a> {
    pub fn new(scores: &'a [u32]) -> Self {
        Self { scores }
    }

    pub fn scores(&self) -> &[u32] {
        self.scores
    }

    pub fn latest(&self) -> Option<u32> {
        self.scores.last().copied()
    }

    pub fn personal_best(&self) -> Option<u32> {
        self.scores.iter().max().copied()
    }

    pub fn personal_top_three(&self) -> Vec<u32> {
        let mut res;

        if self.scores.len() <= 3 {
            res = self.scores.to_vec();
            res.sort_unstable_by(|a, b| b.cmp(a));
            return res;
        }

        res = vec![0; 3];
        for &score in self.scores.iter() {
            for i in 0..res.len() {
                if score > res[i] {
                    res[i..].rotate_right(1);
                    res[i] = score;
                    break;
                }
            }
        }

        res
    }
}
