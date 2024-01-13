pub struct PascalsTriangle {
    row_count: u32,
}

impl PascalsTriangle {
    pub fn new(row_count: u32) -> Self {
        PascalsTriangle { row_count }
    }

    pub fn rows(&self) -> Vec<Vec<u32>> {
        let mut res: Vec<Vec<u32>> = vec![];

        for i in 0..self.row_count {
            if i == 0 {
                res.push(vec![1]);
                continue;
            }

            let mut a = vec![];
            for j in 0..=i {
                if j == 0 || j == i {
                    a.push(1);
                    continue;
                }

                a.push(res.last().unwrap()[j as usize] + res.last().unwrap()[(j - 1) as usize]);
            }

            res.push(a);
        }
        res
    }
}
