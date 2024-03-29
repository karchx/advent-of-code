use std::{fs, process::Command};

pub fn read(day: u32, year: u32, debug: Option<bool>) -> String {
    let root_dir = top_folder();
    let path: String;

    if debug.unwrap_or(false) {
        path = format!("{}/{}/inputs/day{:02}-test.txt", root_dir, year, day);
    } else {
        path = format!("{}/{}/inputs/day{:02}.txt", root_dir, year, day);
    }
    fs::read_to_string(path).expect("Unable to read file")
}

fn top_folder() -> String {
    let system_command = "git";

    let output = Command::new(system_command)
        .arg("rev-parse")
        .arg("--show-toplevel")
        .output()
        .expect("Failed git command");

    let mut output_s = String::from_utf8_lossy(&output.stdout).to_string();
    let len_output = output_s.len();
    output_s.truncate(len_output - 1); // remove \n
    output_s
}
