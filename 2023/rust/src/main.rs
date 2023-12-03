mod days;
mod utils;
mod solution;

use days::day01;
use solution::Solution;

use crate::utils::read;

fn main() {
    let raw_input = read(1, 2023, None); 

    day01::Day1.run(raw_input);
}
