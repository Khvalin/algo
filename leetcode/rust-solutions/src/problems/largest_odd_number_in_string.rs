struct Solution;

impl Solution {
    pub fn largest_odd_number(num: String) -> String {
        let zero = '0' as u8;
        let mut mid = 0;
        let mut chars_iter = num.as_bytes().iter().rev();
        for i in 0..num.len() {
            let ch = chars_iter.next().unwrap();
            if (*ch - zero) % 2 == 1 {
                mid = num.len() - i;
                break;
            }
        }

        (&num[..mid]).to_owned()
    }
}

#[cfg(test)]
mod tests {
    // Note this useful idiom: importing names from outer (for mod tests) scope.
    use super::*;

    #[test]
    fn largest_odd_number() {
        assert_eq!(
            Solution::largest_odd_number("123131231231".to_owned()),
            "123131231231".to_owned()
        );
        assert_eq!(
            Solution::largest_odd_number("12313123123100".to_owned()),
            "123131231231".to_owned()
        );
        assert_eq!(
            Solution::largest_odd_number("666".to_owned()),
            "".to_owned()
        );
    }
}
