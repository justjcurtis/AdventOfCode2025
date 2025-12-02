#!/bin/bash

# Define the README file
README_FILE="README.md"

# Generate the new content
NEW_CONTENT=$(go run . -min -r -n=100)

# Find the start line index (line containing "| - | -------------------- |")
START_LINE=$(grep -n "| - | -------------------- |" "$README_FILE" | cut -d: -f1)

if [ -z "$START_LINE" ]; then
    echo "Start line not found. Exiting."
    exit 1
fi

# Find the next empty line after START_LINE
END_LINE=$(awk "NR > $START_LINE && NF == 0 {print NR; exit}" "$README_FILE")

if [ -z "$END_LINE" ]; then
    echo "End line (empty line) not found. Exiting."
    exit 1
fi

# Create a new file by deleting lines between START_LINE and END_LINE (exclusive)
{
    head -n "$START_LINE" "$README_FILE" # Keep lines up to START_LINE
    echo "$NEW_CONTENT"                 # Insert new content
    tail -n +"$END_LINE" "$README_FILE" # Keep lines from END_LINE onward
} > "${README_FILE}.tmp"

# Replace the original file with the updated file
mv "${README_FILE}.tmp" "$README_FILE"

echo "Content updated between lines $START_LINE and $END_LINE."
echo "New content:"
echo "$NEW_CONTENT"
