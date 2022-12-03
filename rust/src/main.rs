pub mod vowel_count;
pub mod split_strings;

fn main() {

    let cs: Vec<String> =
        "holamundo".chars()
        .collect::<Vec<char>>()
        .chunks(2)
        .map(|x| x.iter().collect())
        .map(|x:String| if x.len() == 1 {x + "_"} else {x})
        .collect();

    dbg!(cs);
}
