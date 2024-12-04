# AOC 2024 attempts by Lotus Engineer (Ryan Burke)

[![xc compatible](https://xcfile.dev/badge.svg)](https://xcfile.dev)

Here it is, some of my hackiest code I've no doubt, Once i'm caught up i'll probably start posting along on [Bluesky](https://bsky.app/profile/lotus-engineer.com) if you want to follow along there, see which days annoyed me the most.

Fresh back from a month off recovering from surgery so i'm a bit rusty, as a result for my first pass i'll probably use a language i'm slightly more familiar with, Go.

If i find myself with the time or interest I may attempt a second pass in a more "exotic" language shall we say, like Zig or Elixir.

## Challenge Micro Blogs

### Go

#### Day 1

Okay nice, fairly simple one to start me off getting back into the swing of solving puzzles with code and the nuances of Go (Day job has me on Python and... Typescript). Probably forgetting some shorthands and over handling errors for a little hacky script (My brain is in production mode) but happy enough with my solution, review from a friend also taught me about go's embed directive for accessing the file, cool one to play with in future.

## Tasks

### run_go_day

Runs the go script for the day provided, for example to run the second day you'd run, `xc run_go_day 2`

```sh
cd ./go_attempts/day_$1/ && go run .
```

### log_go_day

Runs the go script for the day provided, for example to run the second day you'd run, `xc run_go_day 2`

```sh
rm ./log.txt && cd ./go_attempts/day_$1/ && go run . > ../../log.txt
```

### generate_go_day

Runs the go script for the day provided, for example to run the second day you'd run, `xc run_go_day 2`

```sh
mkdir ./go_attempts/day_$1/ && cd ./go_attempts/day_$1/ && touch main.go test.txt input.txt && go mod init github.com/LotusEngineer/advent_of_code && go mod tidy
```
