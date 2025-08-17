#!/bin/bash
set -e

# Analyze files in a directory and generate statistics
DIR=${1:-.}
echo "Analyzing files in: ${DIR}"

echo "=== File Count by Type ==="
find "${DIR}" -type f -name "*.*" \
| while read file; do
    echo "${file##*.}"
done \
| sort | uniq -c | sort -nr

echo ""
echo "=== Largest Files ==="
find "${DIR}" -type f -exec ls -la {} \; \
| awk '{print $5 "\t" $9}' \
| sort -nr | head -10

echo ""
echo "=== Total Size ==="
find "${DIR}" -type f -exec ls -la {} \; \
| awk '{sum += $5} END {print "Total: " sum " bytes"}'
