// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn range_sum_bst(root: Option<Rc<RefCell<TreeNode>>>, low: i32, high: i32) -> i32 {
        fn walk(node: &Option<Rc<RefCell<TreeNode>>>, low: i32, high: i32) -> i32 {
            let mut res = 0;
            if let Some(v) = node {
                let v = v.borrow();

                let val = v.val;

                if val >= low && val <= high {
                    res += val;
                }

                res += walk(&v.left, low, high);
                res += walk(&v.right, low, high);
            }

            res
        }

        let mut res = 0;

        if let Some(v) = root {
            res = walk(&Some(v), low, high);
        }

        res
    }
}
