use std::{env, fs};

use rand::{thread_rng, Rng};
use serde_derive::Deserialize;

#[derive(Deserialize)]
struct Comment {
    text: String,
}

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut exe_path = env::current_exe()?;
    exe_path.pop();
    exe_path.push("comments.json");
    let content = fs::read_to_string(exe_path)?;

    let comments: Vec<Comment> = serde_json::from_str(&content)?;

    if comments.is_empty() {
        return Err("No comments found".into());
    }

    let mut rng = thread_rng();
    let selected_comment = &comments[rng.gen_range(0..comments.len())];

    println!("{}", selected_comment.text);

    return Ok(());
}
