use itertools::Itertools;
use std::cmp::Ordering;
use core::iter::zip;

#[allow(dead_code)]
fn power(a: &[u32]) -> Vec<Vec<u32>> {
    let size:usize = a.len();
    let mut salida:Vec<Vec<u32>>= vec![];
    for i in 0..2_i32.pow(size as u32){
        let mut v = i;
        let mut comb:Vec<u32> = vec![];
        for e in 0..size{
            if v%2 == 1 {
                comb.push(a[e]);
            } 
            v = v / 2;
        }
        salida.push(comb)
    }
    return salida;
}

#[allow(dead_code)]
fn f(a: &[u32], b: &[u32]) -> Ordering {
    for (x, y) in zip(a, b) {
        if x != y {
            return x.cmp(&y)
        }
    }
    return a.len().cmp(&b.len())
}

#[allow(dead_code)]
fn dotest(a: &[u32], expected: &[Vec<u32>]) {
    let actual = power(a);
    assert!(
        actual
            .iter()
            .map(|x| x.iter().sorted().copied().collect::<Vec<_>>())
            .collect::<Vec<_>>()
            .iter()
            .sorted_by(|a, b| f(&a, &b))
            .eq(expected
                .iter()
                .map(|x| x.iter().sorted().copied().collect::<Vec<_>>())
                .collect::<Vec<_>>()
                .iter()
                .sorted_by(|a, b| f(a, b))
            ),
        "With a = {a:?}\nExpected {expected:?} but got {actual:?}")
}

#[test]
fn fixed_tests() {
    dotest(&[1, 2, 3], &[vec![], vec![1], vec![2], vec![1, 2], vec![3], vec![1, 3], vec![2, 3], vec![1, 2, 3]]);
}

#[test]
fn empty_test(){
    dotest(&[], &[vec![]])
}

#[test]
fn single_element_test(){
    dotest(&[1], &[vec![], vec![1]])
}
