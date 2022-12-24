use std::cmp::max;

pub fn run(speed: u32, time: u32) -> u32 {
    if speed <= 1 || time <= 2 {
        return (time + 1) * speed;
    } else {
        let b1 = 3 * speed - 1 + run(speed - 1, time - 2);
        let b2 = speed + run(speed, time - 1);
        println!("run({}, {}) ({}, {}): {} <-> ({}, {}): {} => {}",
                 speed, time,                   // llamada
                 speed-1, time-2, b1,           // Speed
                 speed, time-1, b2,             // Run
                 if b1 > b2 {"S"} else {"R"});  // Resultado
        return max(
            b1,
            b2,
        );
    }
}

pub fn distance(run: &str, speed: u32) -> u32 {
    let mut actual_speed = speed;
    let mut total_distance = 0;
    if speed == 0 {
        return 0;
    };

    for i in run.chars() {
        match i {
            'R' => total_distance = total_distance + actual_speed,
            'S' => {
                total_distance = total_distance + 2 * actual_speed;
                actual_speed -= 1
            }
            _ => return 0,
        }
    }
    return total_distance;
}

#[test]
fn test_zeros() {
    assert_eq!(run(1, 1), 2);
    assert_eq!(run(0, 1), 0);
    assert_eq!(run(0, 2), 0);
}

#[test]
fn newbie() {
    assert_eq!(run(2, 3), 8);
}
#[test]
fn not_so_fast() {
    assert_eq!(run(1, 1), 2);
    assert_eq!(run(2, 1), 4);
}

#[test]
fn test_initials() {
    assert_eq!(run(2, 1), 4);
    assert_eq!(run(4, 1), 8);
    assert_eq!(run(1, 2), 3);
    assert_eq!(run(3, 2), 9);
    assert_eq!(run(4, 2), 12);
    assert_eq!(run(5, 2), 15);
}

#[test]
fn test_table() {
    assert_eq!(run(1, 4), 5);
    assert_eq!(run(2, 3), 8);
    assert_eq!(run(3, 3), 12);
    assert_eq!(run(4, 3), 17);
    assert_eq!(run(2, 4), 10);
    assert_eq!(run(5, 4), 27);
    assert_eq!(run(5, 3), 22)
}

#[test]
fn twos_company() {
    assert_eq!(run(1, 2), 3);
    assert_eq!(run(2, 2), 6);
}

#[test]
#[ignore = "not ready yet"]
fn long_trip() {
    assert_eq!(run(49, 50), 2875);
}
#[test]
#[ignore = "not ready yet"]
fn no_way_home() {
    assert_eq!(run(829, 135), 161453);
}
