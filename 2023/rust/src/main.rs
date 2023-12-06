mod days;
mod utils;
mod solution;

use days::{day01, day02, day03};
use solution::Solution;

use crate::utils::read;

fn main() {
    let raw_input_1 = read(1, 2023, None); 
    let raw_input_2 = read(2, 2023, None); 
    let raw_input_3 = read(3, 2023, None);

   // day01::Day1.run(raw_input_1);
   // day02::Day2.run(raw_input_2);
    day03::Day3.run(raw_input_3);
}
