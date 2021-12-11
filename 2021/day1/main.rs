use std::io;
use std::io::BufRead;

fn main() {
    let mut increased = 0;
    let mut previous: Option<u32> = None;

    let stdin = io::stdin();

    let values: Vec<u32> = stdin
        .lock()
        .lines()
        .map(|x| x.unwrap().parse::<u32>().unwrap())
        .collect();

    for value in &values {
        if let Some(p) = previous {
            if p <= *value {
                increased += 1;
            }
        }
        previous = Some(*value);
    }

    let size = values.len();

    let mut part_b = 0;
    for i in 3..size {
        let a = values[i - 3] + values[i - 2] + values[i - 1];
        let b = values[i - 2] + values[i - 1] + values[i];
        if b > a {
            part_b += 1;
        }
    }

    println!("Part A: {}", increased);
    println!("Part B: {}", part_b);
}
