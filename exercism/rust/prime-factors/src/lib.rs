pub fn factors(mut n: u64) -> Vec<u64> {
    let mut result = vec![];
    let mut p = 2;

    while n > 1 {
        if n % p == 0 {
            result.push(p);
            n /= p;
            continue;
        }

        p += 1;
    }

    result
}
