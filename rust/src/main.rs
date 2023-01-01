use std::cmp::{max, min};


pub mod rgb_conversion;

fn main() {
    let x =-300;
    println!("{:02X?}", max(min(x, 255), 0));
}
