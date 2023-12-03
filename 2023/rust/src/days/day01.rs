use std::fmt::Display;

use crate::solution::Solution;

pub struct Day1;

impl Day1 {
    pub fn trebuchet_partone<T: Display>(&self, input: T) -> u32 {
        input
            .to_string()
            .lines()
            .map(|line| {
                let digits = line
                    .chars()
                    .filter_map(|c| c.to_digit(10))
                    .collect::<Vec<u32>>();
                digits.first().unwrap() * 10 + digits.last().unwrap()
            })
            .sum()
    }

    pub fn trebuchet_parttwo<T: Display>(&self, input: T) -> u32 {
        let mut input = input.to_string();
        let map_numbers = [
            ("one", "1"),
            ("two", "2"),
            ("three", "3"),
            ("four", "4"),
            ("five", "5"),
            ("six", "6"),
            ("seven", "7"),
            ("eight", "8"),
            ("nine", "9"),
        ];

        for (key, val) in map_numbers {
            input = input.replace(key, format!("{key}{val}{key}").as_str());
        }
        self.trebuchet_partone(input)
    }
}

impl Solution for Day1 {
    const NAME: &'static str = "Trebuchet?!";

    fn run(&self, input: String) {
        let p1 = self.trebuchet_partone(&input);
        let p2 = self.trebuchet_parttwo(&input);

        println!("{}", self.name().to_uppercase());
        println!("Part 1 {p1}");
        println!("Part 2 {p2}");
    }
}
