pub fn trebuchet_partone(data: Vec<&str>) {
    let mut sum = 0;
    for ele in data.to_vec() {
       let digits: Vec<i64> = ele
           .to_string()
           .chars()
           .filter_map(|x| x.to_digit(10).map(|x| x as i64))
           .collect();
       
       let decimal_digit = digits[0] * 10 + digits[digits.len() - 1];
       sum += decimal_digit;
    }

    println!("{}", sum);
}

