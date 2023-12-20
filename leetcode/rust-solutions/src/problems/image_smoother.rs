struct Solution;

impl Solution {
    pub fn image_smoother(img: Vec<Vec<i32>>) -> Vec<Vec<i32>> {
        let mut res = img.clone();
        for i in 0..img.len() {
            for j in 0..img[i].len() {
                let mut s = 0;
                let mut c = 0;

                for k in (i as i32 - 1)..=(i as i32 + 1) {
                    for l in (j as i32 - 1)..=(j as i32 + 1) {
                        if k < 0 || l < 0 || k >= (img.len() as i32) || l >= (img[i].len() as i32) {
                            continue;
                        }

                        s += img[k as usize][l as usize];
                        c += 1;
                    }
                }

                res[i][j] = s / c;
            }
        }

        res
    }
}
