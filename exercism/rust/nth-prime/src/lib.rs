use num_integer::Roots;

pub fn nth(n: usize) -> u32 {
    let mut primes = vec![2];
    let mut cur = 3;
    while (&primes).len() <= n {
        let mut is_prime = true;
        let max = cur.sqrt();
        for p in &primes {
            if cur % *p == 0 {
                is_prime = false;
                break;
            }
            if *p > max {
                break;
            }
        }

        if is_prime {
            primes.push(cur);
        }
        cur += 2;
    }

    primes[n]
}
