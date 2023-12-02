mod days;
mod utils;

use crate::utils::read;

fn main() {
    let data = read(1, 2023);

    let lines = data
        .lines()
        .collect::<Vec<_>>();

    for ele in lines {
        println!("{}", ele);
    }
}
