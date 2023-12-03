use std::fmt::Display;

use crate::solution::Solution;

pub struct Day2;

impl Day2 {
    pub fn part_one<T: Display>(&self, input: T) -> u32 {
        input
            .to_string()
            .lines()
            .map(|line| {
                let (id, plays) = line.split_once(":").unwrap();

                let game_id = id.trim_start_matches("Game ").parse::<u32>().unwrap();

                if plays.split(";").all(|play| {
                    let mut red = 0;
                    let mut blue = 0;
                    let mut green = 0;

                    for color in play.splitn(3, ",") {
                        let (num, name) = color.trim().split_once(" ").unwrap();
                        let num = num.parse::<u32>().unwrap();

                        match name {
                            "red" => red += num,
                            "green" => green += num,
                            "blue" => blue += num,
                            _ => (),
                        }
                    }

                    red <= 12 && green <= 13 && blue <= 14
                }) {
                    game_id
                } else {
                    0
                }
            })
            .sum()
    }

    pub fn part_two<T: Display>(&self, input: T) -> u32 {
        input
            .to_string()
            .lines()
            .map(|line| {
                let (_, plays) = line.split_once(":").unwrap();

                let mut red = 0;
                let mut green = 0;
                let mut blue = 0;

                for colors in plays.replace(";", ",").split(",") {
                    let (num, name) = colors.trim().split_once(" ").unwrap();

                    let num = num.parse::<u32>().unwrap();
                    match name {
                        "red" if num > red => red = num,
                        "green" if num > green => green = num,
                        "blue" if num > blue => blue = num,
                        _ => (),
                    }
                }
                red * green * blue
            })
            .sum()
    }
}

impl Solution for Day2 {
    const NAME: &'static str = "Cube Conundrum";

    fn run(&self, input: String) {
        let p1 = self.part_one(&input);
        let p2 = self.part_two(&input);
        println!("{}", self.name().to_uppercase());
        println!("Part1 {}", p1);
        println!("Part2 {}", p2);
    }
}
