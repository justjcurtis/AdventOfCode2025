#!/bin/bash

# Get .env vars
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

# Get the puzzleInput directory path
PUZZLE_DIR="./puzzleInput"
SOLUTION_DIR="./solutions"
MAIN_FILE="./main.go"

# Check if puzzleInput directory exists
if [ ! -d "$PUZZLE_DIR" ]; then
    echo "Error: $PUZZLE_DIR directory does not exist"
    exit 1
fi

# Find all day_x.txt files and extract the number x
echo "Scanning for day_x.txt files in $PUZZLE_DIR..."

# Create an array to store found numbers
declare -a numbers

# Loop through all files matching the pattern
for file in "$PUZZLE_DIR"/day_*.txt; do
    if [ -f "$file" ]; then
        # Extract the number from filename using regex
        if [[ $(basename "$file") =~ ^day_([0-9]+)\.txt$ ]]; then
            number=${BASH_REMATCH[1]}
            numbers+=($number)
        fi
    fi
done

# If no files found, exit
if [ ${#numbers[@]} -eq 0 ]; then
    echo "No day_x.txt files found in $PUZZLE_DIR"
    exit 0
fi

# Find the maximum number to determine what to create
max_number=0
for num in "${numbers[@]}"; do
    if [ "$num" -gt "$max_number" ]; then
        max_number=$num
    fi
done

# Create the new files with number + 1
DAY=$((max_number + 1))

# Update main.go
DAY_ENTRY="        {${DAY}, solutions.Day${DAY}},"

awk -v new_entry="$DAY_ENTRY" '
BEGIN {added=0}
/var SOLUTIONS = \[\]solution\s*{/ {
    print
    in_block=1
    next
}
in_block && /^\}/ {
    print new_entry
    added=1
    in_block=0
}
{print}
END {
    if (!added) {
        print "// WARNING: Could not add new day to SOLUTIONS"
    }
}
' "$MAIN_FILE" > "${MAIN_FILE}.tmp" && mv "${MAIN_FILE}.tmp" "$MAIN_FILE"

echo "Added Day ${DAY} to SOLUTIONS list"

# Create the new day file
touch "$PUZZLE_DIR/day_${DAY}.txt"
echo "Created day_${DAY}.txt"

# Fetch and save the puzzle input
echo "Fetching puzzle input for Day $DAY..."
URL="https://adventofcode.com/${YEAR}/day/${DAY}/input"

curl "$URL" \
    --cookie "session=${COOKIE}" \
    -A "bash script for personal use" \
    -s > "$PUZZLE_DIR/day_${DAY}.txt"
echo "Fetched puzzle input for Day $DAY"

# Create the new test file
touch "$PUZZLE_DIR/test_${DAY}.txt"
echo "Created test_${DAY}.txt"

# Create a solution file template
SOLUTION_FILE="$SOLUTION_DIR/day${DAY}.go"
mkdir -p "$SOLUTION_DIR"
cat <<EOL > "$SOLUTION_FILE"
package solutions

func Day${DAY}(input []string) []string {
    return []string{""}
}
EOL
echo "Created day${DAY}.go"

# Create a test file template
TEST_FILE="$SOLUTION_DIR/day${DAY}_test.go"
cat <<EOL > "$TEST_FILE"
package solutions
import (
    "AdventOfCode${YEAR}/utils"
    "testing"
    "reflect"
)

func TestDay${DAY}(t *testing.T) {
    input := utils.GetTestInput(${DAY})
    expected := []string{""}
    actual := Day${DAY}(input)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v but was %v", expected, actual)
	}
}
EOL
echo "Created day${DAY}_test.go"

