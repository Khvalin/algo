pub fn brackets_are_balanced(string: &str) -> bool {
    let mut stack = vec![];

    for ch in string.chars() {
        match ch {
            '(' | '[' | '{' => {
                stack.push(ch);
                continue;
            }
            '}' | ']' | ')' => {
                let last = stack.pop();
                if last.is_none() {
                    return false;
                }

                if (last.unwrap() == '(' && ch != ')')
                    || (last.unwrap() == '[' && ch != ']')
                    || (last.unwrap() == '{' && ch != '}')
                {
                    return false;
                }
            }
            _ => continue,
        }
    }

    stack.is_empty()
}
