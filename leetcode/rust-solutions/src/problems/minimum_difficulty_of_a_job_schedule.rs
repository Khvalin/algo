struct Solution;

use std::collections::HashMap;

impl Solution {
    pub fn solve(
        memo: &mut Vec<Vec<Option<i32>>>,
        job_difficulty: &Vec<i32>,
        ind: usize,
        d: i32,
    ) -> i32 {
        if d == 0 && ind >= job_difficulty.len() {
            return 0;
        }
        if ((job_difficulty.len() - ind) as i32) < d || d < 0 {
            // [1,2,3], d = 4
            return i32::MAX;
        }

        if memo[ind][d as usize].is_some() {
            return memo[ind][d as usize].unwrap();
        }

        let mut res = i32::MAX;
        let mut max = i32::MIN;
        for i in ind..job_difficulty.len() {
            max = max.max(job_difficulty[i]);
            let t = Solution::solve(memo, job_difficulty, i + 1, d - 1);
            if t < i32::MAX {
                res = res.min(t + max);
            }
        }

        memo[ind][d as usize] = Some(res);

        res
    }

    pub fn min_difficulty(job_difficulty: Vec<i32>, d: i32) -> i32 {
        let mut memo = vec![vec![None; d as usize + 1]; job_difficulty.len() + 1];
        let res = Solution::solve(&mut memo, &job_difficulty, 0, d);
        if res == i32::MAX {
            return -1;
        }

        return res;
    }
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn min_difficulty() {
        assert_eq!(Solution::min_difficulty(vec![6, 5, 4, 3, 2, 1], 2), 7);
        let v = vec![
            42, 35, 957, 671, 87, 222, 524, 5, 280, 878, 242, 398, 610, 984, 649, 513, 563, 997,
            786, 223, 413, 961, 208, 189, 519, 606, 504, 406, 994, 475, 940, 476, 762, 283, 218,
            404, 715, 755, 689, 457, 20, 403, 749, 68, 17, 763, 78, 695, 445, 338, 643, 400, 615,
            29, 866, 861, 103, 665, 743, 361, 798, 236, 255, 15, 326, 124, 14, 588, 711, 992, 501,
            731, 538, 619, 765, 8, 477, 854, 243, 95, 627, 480, 505, 800, 818, 722, 194, 717, 305,
            810, 497, 686, 704, 322, 532, 22, 234, 290, 885, 416, 155, 428, 161, 315, 152, 441,
            730, 926, 175, 536, 646, 922, 951, 101, 107, 233, 61, 735, 669, 512, 28, 569, 447, 982,
            916, 321, 1000, 754, 483, 482, 811, 654, 47, 344, 772, 978, 467, 308, 472, 269, 0, 252,
            131, 834, 303, 945, 469, 785, 537, 188, 449, 675, 528, 733, 271, 541, 822, 328, 904,
            876, 889, 55, 16, 853, 154, 767, 573, 925, 279, 697, 525, 431, 375, 958, 836, 911, 166,
            965, 523, 709, 923, 587, 603, 226, 354, 641, 620, 316, 110, 292, 529, 943, 1, 151, 737,
            959, 27, 56, 353, 681, 26, 677, 337, 723, 12, 914, 955, 134, 370, 260, 490, 684, 364,
            618, 232, 870, 985, 196, 225, 359, 129, 58, 341, 67, 494, 757, 229, 323, 256, 21, 783,
            692, 642, 37, 909, 248, 81, 345, 425, 478, 651, 309, 435, 10, 534, 450, 576, 144, 201,
            496, 267, 609, 11, 454, 531, 966, 156, 176, 190, 542, 856, 365, 983, 947, 758, 950,
            388, 820, 867, 833, 605, 741, 34, 663, 884, 65, 628, 969, 864, 664, 718, 805, 891, 657,
            863, 960, 518, 71, 300, 756, 613, 667, 228, 274, 971, 970, 552, 556, 648, 251,
        ];
        assert_eq!(Solution::min_difficulty(v, 10), 3044);
    }
}
