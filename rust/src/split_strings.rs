#[allow(unused_variables)]

/*
La primera solución es más fea pero no tiene collect al principio
lo que haría que funcionara mejor en caso una entrada absurdamente
grande.

Pero la segunda es más elegante.
*/

#[allow(dead_code)]
fn solution_2(s: &str) -> Vec<String> {
    //let mut res = vec!["".to_string()];
    
    let mut res:Vec<String>=Vec::new();

    let mut cs = s.chars();
    loop {
        let x = cs.next();
        if x.is_none() {
            break;
        }
        let y: Option<char> = cs.next();
        if y.is_none(){
            res.push(format!("{}{}", x.unwrap(),'_'));
            break
        }
        res.push(format!("{}{}", x.unwrap(), y.unwrap()));
    }
    res
}

#[allow(dead_code)]
fn solution(s: &str) -> Vec<String>{
    let cs: Vec<String> =s.chars()
        .collect::<Vec<char>>()
        .chunks(2)
        .map(|x| x.iter().collect())
        .map(|x:String| if x.len() == 1 {x + "_"} else {x})
        .collect();

    cs
}

#[test]
fn basic() {
    assert_eq!(solution("abcdef"), ["ab", "cd", "ef"]);
    assert_eq!(solution("abcdefg"), ["ab", "cd", "ef", "g_"]);
    assert_eq!(solution(""), [] as [&str; 0]);
}
