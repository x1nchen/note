use clap::Parser;

#[derive(Parser, Debug)]
#[clap(
    version = "1.0",
    author = "ccx",
)]
struct Opts {
    #[clap(subcommand)]
    subcmd: SubCommand,
}

#[derive(Parser, Debug)]
enum SubCommand {
    Get(Get),
    Post(Post),
}

// 子命令
#[derive(Parser, Debug)]
struct Get {
    url: String,
}

#[derive(Parser, Debug)]
struct Post {
    url: String,
    body: Vec<String>,
}

fn main() {
    let opts = Opts::parse();
    println!("{:?}", opts);
}
