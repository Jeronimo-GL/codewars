

#[allow(dead_code)]
#[allow(unused_variables)]
fn remove_parentheses(s: &str) -> String {
    let mut sc = s.to_string();
    let mut to_remove:Vec<usize> = Vec::new();
    
    for (pos, c) in s.chars().enumerate() {
        match c {
            '(' => {
                to_remove.push(pos);
            },
            ')' => {
                if let Some(start) = to_remove.pop() {
                    sc.replace_range(start .. pos+1, &"#".repeat((pos-start)+1));
                }
            },
            _ => (),
        }
    }
    return sc.replace("#", "");
    
}


#[test]
fn sample_tests() {
    assert_eq!(remove_parentheses("example(unwanted thing)example"), "exampleexample");
    assert_eq!(remove_parentheses("example (unwanted thing) example"), "example  example");
    assert_eq!(remove_parentheses("a (bc d)e"), "a e");
    assert_eq!(remove_parentheses("a(b(c))"), "a");
    assert_eq!(remove_parentheses("hello example (words(more words) here) something"), "hello example  something");
    assert_eq!(remove_parentheses("(first group) (second group) (third group)"), "  ");
    assert_eq!(remove_parentheses("(todo)"), "");
    assert_eq!(remove_parentheses(""), "");
    assert_eq!(remove_parentheses("nada"), "nada");
}
