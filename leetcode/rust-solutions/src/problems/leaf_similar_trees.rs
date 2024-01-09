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

struct Solution;

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn leaf_similar(
        root1: Option<Rc<RefCell<TreeNode>>>,
        root2: Option<Rc<RefCell<TreeNode>>>,
    ) -> bool {
        fn walk(node: &Option<Rc<RefCell<TreeNode>>>, vals: &mut Vec<i32>) {
            if let Some(v) = node {
                let v = v.borrow();

                match (&v.left, &v.right) {
                    (None, None) => vals.push(v.val),
                    _ => {
                        walk(&v.left, vals);
                        walk(&v.right, vals);
                    }
                }
            }
        }

        let mut leaves1: Vec<i32> = Vec::with_capacity(200);
        if let Some(v) = root1 {
            walk(&Some(v), &mut leaves1);
        }

        let mut leaves2: Vec<i32> = Vec::with_capacity(200);
        if let Some(v) = root2 {
            walk(&Some(v), &mut leaves2);
        }

        leaves1 == leaves2
    }
}
