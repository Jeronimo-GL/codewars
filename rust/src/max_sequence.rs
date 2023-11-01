use std::cmp::max;
// https://www.codewars.com/kata/54521e9ec8e60bc4de000d6c/train/rust

#[allow(dead_code)]
fn max_sequence(seq: &[i32]) -> i32 {
    if seq.len() == 0 {
        return 0;
    }
    let mut maxes = vec![0; seq.len()];
    for i in 0..(seq.len()) {
        dbg!(&i);
        for j in 0..(seq.len() - i) {
            dbg!(&j);
            dbg!(&seq[i..i + j + 1]);
            let local_sum: i32 = seq[i..i + j + 1].iter().sum();
            dbg!(&local_sum);
            maxes[j] = max(maxes[j], local_sum);
        }
    }
    dbg!(&maxes);
    return *maxes.iter().max().unwrap_or(&0);
}

#[cfg(test)]
mod tests {
    use super::max_sequence;

    #[test]
    fn sample_tests() {
        assert_eq!(max_sequence(&[11]), 11);
        assert_eq!(max_sequence(&[-32]), 0);
        assert_eq!(max_sequence(&[-2, 1, -3, 4, -1, 2, 1, -5, 4]), 6);
    }

    #[test]
    fn negative_sums() {
        assert_eq!(max_sequence(&[-1]), 0);
        assert_eq!(max_sequence(&[-1, -1, -3]), 0);
    }

    #[test]
    fn empty_list() {
        assert_eq!(max_sequence(&[]), 0);
    }

    #[test]
    fn single_value_lists() {
        assert_eq!(max_sequence(&[2]), 2);
        assert_eq!(max_sequence(&[0]), 0);
        assert_eq!(max_sequence(&[1000]), 1000);
    }

    #[test]
    fn two_values() {
        assert_eq!(max_sequence(&[2, 2]), 4);
        assert_eq!(max_sequence(&[-1, 3]), 3);
        assert_eq!(max_sequence(&[3, -1]), 3);
    }
}
