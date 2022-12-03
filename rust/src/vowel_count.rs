

#[allow(dead_code)]
fn get_count(string: &str) -> usize {
  let vowels = vec!['a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'];
  let vowels_count: usize = string.chars().filter(|l| vowels.contains(l)).count();
  
  vowels_count
}


#[test]
fn my_tests() {
  assert_eq!(get_count("abracadabra"), 5);
}


#[test]
fn test_empty_string(){
    assert_eq!(get_count(""), 0);
}

#[test]
fn test_caps(){
    assert_eq!(get_count("AA"), 2);
}
