#!/bin/bash
set -e

# Process log files to extract errors and warnings
for file in logs/*.log; do
    if [[ -f "$file" ]]; then
        echo "Processing $file"
        grep -i "error\|warning" "$file" | \
        while read line; do
            timestamp=$(echo "$line" | cut -d' ' -f1)
            level=$(echo "$line" | cut -d' ' -f2)
            echo "$timestamp,$level" >> results.csv
        done
    fi
done

echo "Results written to results.csv"
