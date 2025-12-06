#!/bin/sh

README_FILE="README.md"
NEW_CONTENT=$(go run . -min -r -n=1000)

START_LINE=$(grep -n "| - | -------------------- |" "$README_FILE" | cut -d: -f1)
[ -z "$START_LINE" ] && echo "Start line not found" && exit 1

END_LINE=$(awk "NR > $START_LINE && NF == 0 {print NR; exit}" "$README_FILE")
[ -z "$END_LINE" ] && echo "End line not found" && exit 1

EXISTING_TABLE=$(sed -n "$((START_LINE+1)),$((END_LINE-1))p" "$README_FILE")

UPDATED_TABLE=""
TOTAL=0

while IFS= read -r line; do
    case "$line" in
        "| Day "*)
            day=$(printf "%s\n" "$line" | sed -E 's/\| *(Day [0-9]+) \|.*/\1/')
            old_time=$(printf "%s\n" "$line" | sed -E 's/.*\| *([0-9]+)µs.*/\1/')

            new_line=$(printf "%s\n" "$NEW_CONTENT" | grep "$day" || true)
            if [ -n "$new_line" ]; then
                new_time=$(printf "%s\n" "$new_line" | sed -E 's/.*\| *([0-9]+)µs.*/\1/')
            else
                new_time=$old_time
            fi

            if [ "$new_time" -lt "$old_time" ]; then
                time="$new_time"
            else
                time="$old_time"
            fi

            UPDATED_TABLE="${UPDATED_TABLE}| $day | ${time}µs |\n"
            TOTAL=$((TOTAL + time))
            ;;
        *)
            # skip totals
            :
            ;;
    esac
done <<EOF
$EXISTING_TABLE
EOF

UPDATED_TABLE="${UPDATED_TABLE}| ------- | ----------------------------- |\n"
UPDATED_TABLE="${UPDATED_TABLE}| **Total** | **${TOTAL}µs** |"

{
    head -n "$START_LINE" "$README_FILE"
    printf "%b" "$UPDATED_TABLE"
    tail -n +"$END_LINE" "$README_FILE"
} > "${README_FILE}.tmp"

mv "${README_FILE}.tmp" "$README_FILE"

git add "$README_FILE"
git commit -m "Update performance table in README.md"
