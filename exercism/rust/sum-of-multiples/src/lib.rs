use std::collections::HashSet;

pub fn sum_of_multiples(limit: u32, factors: &[u32]) -> u32 {
    let mut a = HashSet::new();

    for f in factors {
        if *f <= 0 {
            continue;
        }
        for i in (*f..limit).step_by(*f as usize) {
            a.insert(i);
        }
    }

    a.into_iter().sum()
}
