pub mod max_sequence;

fn main() {
    let mut counter: i32 = 1;
    let base: u64 = 10;
    println!("{:>10}          0 d: {:>3} s: {:>8} {}", counter, 0, 0, 0);
    counter = counter + 1;
    (0..7).for_each(|d| {
        ((base.pow(d) as u64)..(base.pow(d + 1) as u64)).for_each(|s| {
            if !(s > ((base.pow(d) + 2) as u64) && s < ((base.pow(d + 1) - 2) as u64)) {
                println!(
                    "{:>10} log: {:>3} O d: {:>3} s: {:>8} {}",
                    counter,
                    (counter as f64).log10().floor(),
                    d,
                    s,
                    complete_odd(&s.to_string())
                );
            }
            counter = counter + 1;
        });
        ((base.pow(d) as u64)..(base.pow(d + 1) as u64)).for_each(|s| {
            if !(s > ((base.pow(d) + 2) as u64) && s < ((base.pow(d + 1) - 2) as u64)) {
                println!(
                    "{:>10} log: {:>3} E d: {:>3} s: {:>8} {}",
                    counter,
                    (counter as f64).log10().floor(),
                    d,
                    s,
                    complete_even(&s.to_string())
                );
            }
            counter = counter + 1;
        });
    });
}

fn complete_odd(n: &str) -> String {
    let part: String = n.clone()[0..n.len() - 1].chars().rev().collect();

    let mut val = n.clone().to_string();
    val.push_str(&part);
    return val;
}

fn complete_even(n: &str) -> String {
    let part: String = n.clone()[0..n.len()].chars().rev().collect();

    let mut val = n.clone().to_string();
    val.push_str(&part);
    return val;
}
