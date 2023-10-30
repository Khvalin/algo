pub fn build_proverb(list: &[&str]) -> String {
    let mut lines = vec![];
    for i in 1..list.len() {
        lines.push(format!(
            "For want of a {} the {} was lost.",
            list[i - 1],
            list[i]
        ));
    }

    if !list.is_empty() {
        lines.push(format!("And all for the want of a {}.", list[0]));
    }

    lines.join("\n")
}
