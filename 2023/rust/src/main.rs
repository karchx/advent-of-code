mod days;
mod utils;

use days::day01;

use crate::utils::read;

fn main() {
    let data = read(1, 2023, None);

    let lines = data
        .lines()
        .collect::<Vec<_>>();

    day01::trebuchet_partone(lines); 
}
