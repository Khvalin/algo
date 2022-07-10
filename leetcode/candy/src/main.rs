struct Solution {}

impl Solution {
    pub fn candy(ratings: Vec<i32>) -> i32 {
        let mut res = ratings.iter().map(|_| 0).collect::<Vec<_>>();
        let mut indices: Vec<usize> = vec![];

        for (i, n) in ratings.iter().enumerate() {
            let mut f = true;
            if i > 0 && *n > ratings[i - 1] {
                f = false;
            }

            if i < ratings.len() - 1 && *n > ratings[i + 1] {
                f = false;
            }

            if f {
                indices.push(i);
            }
        }

        let mut k = 0;
        let mut cur = 1;
        while k < indices.len() {
            let l = indices.len();

            while k < l {
                let i = indices[k];
                res[i] = cur;

                if i > 0 && ratings[i] < ratings[i - 1] {
                    indices.push(i - 1);
                }

                if i < ratings.len() - 1 && ratings[i] < ratings[i + 1] {
                    indices.push(i + 1);
                }

                k += 1;
            }

            cur += 1;
        }

        return res.iter().sum();
    }
}

fn main() {
    println!("{}", Solution::candy(vec![1, 2, 2]));
    println!("{}", Solution::candy(vec![1, 0, 2]));
}

#[test]
fn test() {}
