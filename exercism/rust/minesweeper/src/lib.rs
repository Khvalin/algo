pub fn annotate(minefield: &[&str]) -> Vec<String> {
    let mut result: Vec<_> = minefield
        .iter()
        .map(|&row| row.as_bytes().to_vec())
        .collect();

    for i in 0..minefield.len() {
        let row = minefield[i].as_bytes();
        for j in 0..row.len() {
            if row[j] != '*' as u8 {
                continue;
            }

            for dx in (-1 as i8)..=1 {
                for dy in (-1 as i8)..=1 {
                    if dx == 0 && dy == 0 {
                        continue;
                    }

                    let x = i as i8 + dx;
                    let y = j as i8 + dy;

                    if x < 0 || y < 0 || x >= minefield.len() as i8 || y >= row.len() as i8 {
                        continue;
                    }

                    let l = x as usize;
                    let k = y as usize;
                    if result[l][k] == '*' as u8 {
                        continue;
                    }

                    if result[l][k] == ' ' as u8 {
                        result[l][k] = '0' as u8;
                    }

                    result[l][k] += 1;
                }
            }
        }
    }

    result
        .iter()
        .map(|row| String::from_utf8(row.to_vec()).unwrap())
        .collect()
}
