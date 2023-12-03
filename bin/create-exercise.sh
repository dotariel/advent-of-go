#!/bin/bash

if [[ ! $# -eq 2 ]]; then
  echo "usage: create-exercise YEAR DAY"
  exit
fi

year=$1
day=$2
dir=${year}/day-${day}

mkdir -p ${dir}

cat << EOF > ${dir}/main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Part 1: %v\n", Part1())
	fmt.Printf("Part 2: %v\n", Part2())
}

func Part1() interface{} {
	return 0
}

func Part2() interface{} {
	return 0
}
EOF


cat << EOF > ${dir}/day-${day}.go
package main
EOF

cat << EOF > ${dir}/day-${day}_test.go
package main
EOF

touch ${dir}/README.md
touch ${dir}/input.txt
