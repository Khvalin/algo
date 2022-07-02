struct Solution;

impl Solution {
    pub fn max_area(h: i32, w: i32, horizontal_cuts: Vec<i32>, vertical_cuts: Vec<i32>) -> i32 {
        fn sort<A, T>(mut array: A) -> A
        where
            A: AsMut<[T]>,
            T: Ord,
        {
            let slice = array.as_mut();
            slice.sort();

            array
        }

        let hc = sort(horizontal_cuts);
        let vc = sort(vertical_cuts);

        let mut hmax = Ord::max(hc[0], h - hc[hc.len() - 1]);
        for i in 1..=hc.len() - 1 {
            hmax = Ord::max(hmax, hc[i] - hc[i - 1])
        }

        let mut vmax = Ord::max(vc[0], w - vc[vc.len() - 1]);
        for i in 1..=vc.len() - 1 {
            vmax = Ord::max(vmax, vc[i] - vc[i - 1])
        }

        return (((vmax as u64) * (hmax as u64)) % 1000000007) as i32;
    }
}

fn main() {
    println!("Hello, world!");
}
