pub fn is_armstrong_number(num: u32) -> bool {
    if num == 0 {
        return true;
    }

    let len = num.ilog10() + 1;

    let mut s = 0;
    let mut t = num;

    while s <= num && t > 0 {
        let p = (t % 10).checked_pow(len);
        if p.is_none() {
            return false;
        }

        let sum = s.checked_add(p.unwrap());
        if sum.is_none() {
            return false;
        }

        s = sum.unwrap();

        t /= 10;
    }

    s == num
}
