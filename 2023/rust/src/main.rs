use crate::utils::read;

mod utils;

fn main() {
    let path = read(1, 2023);
    println!("Path {:?}", path);
}
