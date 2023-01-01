
pub mod powerset;

fn main() {
    let n: u32 = 7;
    for i in 0..4{
        if n&(1<<i) > 0{
            print!("{}", 1);
        } else {
            print!("{}", 0);
        }

    }
}
