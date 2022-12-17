struct Solution;

impl Solution {
    pub fn eval_rpn(tokens: Vec<String>) -> i32 {
        let mut stack = Vec::new();

        for t in tokens {
            let token = t.parse::<i32>();
            if token.is_ok() {
                stack.push(token.unwrap());
                continue;
            }
            let n2 = stack.pop().unwrap();
            let n1 = stack.pop().unwrap();
            let r;
            match t.chars().next().unwrap() {
                '+' => r = n1 + n2,
                '-' => r = n1 - n2,
                '*' => r = n1 * n2,
                '/' => r = n1 / n2,
                _ => todo!(),
            }
            stack.push(r);
        }

        *stack.last().unwrap()
    }
}

fn main() {
    //let v = vec!["2","1","+","3","*"];
    let v = vec![
        "10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+",
    ];
    let input = v.iter().map(|&s| s.to_string()).collect::<Vec<_>>();
    let res = Solution::eval_rpn(input);
    println!("{} == 22", res);
}
