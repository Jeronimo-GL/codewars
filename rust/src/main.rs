pub mod chaser;
pub mod remove_pars;
pub mod mini_string_fuck;

// pub mod split_strings;
// pub mod vowel_count;


fn main() {
    let s: &str = "Hola (con) parentesis() mas ( sds ( ds ) iuy)";
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
    println!("{}", &s);
    let fs = sc.replace("#", "");
    println!("{}", &fs);
}
