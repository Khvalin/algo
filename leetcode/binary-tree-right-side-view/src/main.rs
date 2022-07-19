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


use std::collections::VecDeque;

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn right_side_view(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut  res: Vec<i32> = vec![];
        if root.is_none() {
            return res;
        }

        let mut res = vec![];
        let mut q = VecDeque::new();
        q.push_back(root.unwrap());

        while q.len() > 0 {

            let end = q.len() - 1;
            res.push( q[end].borrow().val.clone() );

            for _ in 0..=end {
                let node = q.pop_front().unwrap();
                if let Some(ref left) = node.clone().borrow().left {
                    q.push_back(left.clone());
                }
                if let Some(ref right) = node.clone().borrow().right {
                    q.push_back(right.clone());
                }
            }

        }

        return res;
    }
}

fn main() {
    println!("Hello, world!");
}
