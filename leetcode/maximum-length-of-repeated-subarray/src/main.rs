struct Solution;

impl Solution {
    pub fn find_length(nums1: Vec<i32>, nums2: Vec<i32>) -> i32 {
        let mut memo: Vec<Vec<i32>> = vec![];
        for _ in nums1.iter() {
            memo.push( vec![0; nums2.len()+1]);
        }

        let mut res = 0;
        for i in (0..nums1.len()).rev() {
            for j in (0..nums2.len()).rev() {
                if nums1[i] != nums2[j] {
                    continue
                }

                let mut v = 1;
                if i < nums1.len() - 1 && j < nums2.len() - 1 && nums1[i + 1] == nums2[j + 1] {
                    v += memo[i+1][j+1]
                }
                memo[i][j] = v;
                res = res.max(v);
            }
        }

        res
    }
}

fn main() {
    let res = Solution::find_length(vec![1,2,3,2,1], vec![3,2,1,4,7]);
    println!("{}", res);
}
