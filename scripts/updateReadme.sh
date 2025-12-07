
README_FILE="README.md"
NEW_CONTENT=$(go run . -min -r -n=1000)

START_LINE=$(grep -n "| - | -------------------- |" "$README_FILE" | cut -d: -f1)
[ -z "$START_LINE" ] && echo "Start line not found" && exit 1

END_LINE=$(awk "NR > $START_LINE && NF == 0 {print NR; exit}" "$README_FILE")
[ -z "$END_LINE" ] && echo "End line not found" && exit 1

EXISTING_TABLE=$(sed -n "$((START_LINE+1)),$((END_LINE-1))p" "$README_FILE")

UPDATED_TABLE=""
TOTAL=0

EXISTING_NUMS=$(printf "%s\n" "$EXISTING_TABLE" | sed -n 's/| *Day \([0-9][0-9]*\) .*/\1/p')
NEW_NUMS=$(printf "%s\n" "$NEW_CONTENT"    | sed -n 's/| *Day \([0-9][0-9]*\) .*/\1/p')

ALL_NUMS=$(
  printf "%s\n%s\n" "$EXISTING_NUMS" "$NEW_NUMS" |
  sed '/^$/d' |
  sort -n |
  uniq
)

if [ -z "$ALL_NUMS" ]; then
  echo "No Day entries found in README or new output. Aborting to avoid wiping table."
  exit 1
fi

extract_time() {
  awk -F'|' -v d="$2" '
    {
      gsub(/^[[:space:]]+|[[:space:]]+$/, "", $2)
      if ($2 == "Day " d) {
        t = $3
        gsub(/[^0-9]/, "", t)
        if (t != "") { print t; exit }
      }
    }
  ' <<EOF
$1
EOF
}

for n in $ALL_NUMS; do
  day="Day $n"

  old_time=$(extract_time "$EXISTING_TABLE" "$n" || true)
  new_time=$(extract_time "$NEW_CONTENT"    "$n" || true)

  if [ -n "$old_time" ] && [ -n "$new_time" ]; then
    if [ "$new_time" -lt "$old_time" ]; then
      time="$new_time"
    else
      time="$old_time"
    fi
  elif [ -n "$old_time" ]; then
    time="$old_time"
  else
    time="$new_time"
  fi

  if [ -z "$time" ]; then
    echo "Warning: no time found for $day; skipping."
    continue
  fi

  UPDATED_TABLE="${UPDATED_TABLE}| $day | ${time}µs |\n"
  TOTAL=$((TOTAL + time))
done

UPDATED_TABLE="${UPDATED_TABLE}| ------- | ----------------------------- |\n"
UPDATED_TABLE="${UPDATED_TABLE}| **Total** | **${TOTAL}µs** |\n"

{
  head -n "$START_LINE" "$README_FILE"
  printf "%b\n" "$UPDATED_TABLE"
  tail -n +"$((END_LINE + 1))" "$README_FILE"
} > "${README_FILE}.tmp"

mv "${README_FILE}.tmp" "$README_FILE"

git add "$README_FILE"
git commit -m "Update performance table in README.md"
