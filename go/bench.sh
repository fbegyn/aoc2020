#!/usr/bin/env bash

days=()

for file in ./bin/*; do
    days+=("${file}")
done

echo $days

hyperfine --warmup 10 --export-csv aoc2020.csv "${days[@]}"
