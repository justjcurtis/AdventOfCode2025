#!/bin/bash

# GET env vars
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

# Find all files and replace 20xx with $YEAR
find "." -type f \( \
    -name "*.md" -o \
    -name "*.go" -o \
    -name "*.txt" -o \
    -name "*.mod" \
\) | while read -r file; do
    perl -pi -e "s/20\d\d/$YEAR/g" "$file"
done

echo "Updated YEAR=$YEAR across project"
