use std::fmt::Display;

use crate::solution::Solution;

pub struct Day3;

impl Day3 {
    fn symbol_adjacent<F, A>(arr: &[Vec<char>], coordinates: A, condition: F) -> Vec<(usize, usize)>
    where
        F: Fn(char) -> bool,
        A: AsRef<[(usize, usize)]>,
    {
        let coordinates = coordinates.as_ref();
        let ncoords = coordinates.len() - 1;

        coordinates
            .iter()
            .enumerate()
            .flat_map(|(i, (row, col))| {
                let mut indices = vec![(*row - 1, *col), (*row + 1, *col)];
                if i == 0 {
                    indices.extend([(*row, *col - 1), (*row - 1, *col - 1), (*row + 1, *col - 1)]);
                }
                if i == ncoords {
                    indices.extend([(*row, *col + 1), (*row - 1, *col + 1), (*row + 1, *col + 1)]);
                }
                indices.into_iter().filter_map(|(y, x)| {
                    arr.get(y)
                        .and_then(|row| row.get(x).and_then(|c| condition(*c).then_some((y, x))))
                })
            })
            .collect()
    }

    pub fn part_one<T: Display>(&self, inp: T) -> u32 {
        let arr = inp
            .to_string()
            .lines()
            .map(|line| line.chars().collect())
            .collect::<Vec<Vec<char>>>();
        let mut total = 0;

        let mut curr_indices = Vec::new();
        let mut curr_num = String::new();
        for (y, row) in arr.iter().enumerate() {
            for (x, chr) in row.iter().enumerate() {
                if chr.is_numeric() {
                    curr_indices.push((y, x));
                    curr_num.push(*chr);
                } else {
                    if !curr_indices.is_empty()
                        && !Self::symbol_adjacent(&arr, &curr_indices, |c| {
                            !c.is_numeric() && c != '.'
                        })
                        .is_empty()
                    {
                        total += curr_num.parse::<u32>().unwrap();
                    }
                    curr_indices.clear();
                    curr_num.clear();
                }
            }
        }
        total
    }
}

impl Solution for Day3 {
    const NAME: &'static str = "Gear Ratios";

    fn run(&self, input: String) {
        let p1 = self.part_one(&input);
        //let p2 = self.part_two(&input);
        println!("{}", self.name().to_uppercase());
        println!("Part1 {}", p1);
        //println!("Part2 {}", p2);
    }
}
