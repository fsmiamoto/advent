#![feature(stdin_forwarders)]
use std::io;
use std::str::FromStr;

fn main() {
    let mut increased = 0;
    let mut previous: Option<u32> = None;
    io::stdin().lines().into_iter().for_each(|l| {
        // ignoring all errors :D
        let line = l.unwrap();
        let current = u32::from_str(String::from(line).as_str()).unwrap();
        if let Some(p) = previous {
            if p <= current {
                increased += 1;
            }
        }
        previous = Some(current)
    });
    println!("{}", increased);
}
