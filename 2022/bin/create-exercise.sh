#!/bin/bash

if [[ ! $# -eq 1 ]]; then
  echo "usage: create-exercise DAY"
  exit
fi

day=$1
dir=day-${day}


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

func Part1() int {
	return 0
}

func Part2() int {
	return 0
}
EOF


cat << EOF > ${dir}/day${day}.go
package main
EOF

cat << EOF > ${dir}/day${day}_test.go
package main
EOF

touch ${dir}/README.md
touch ${dir}/input.txt
