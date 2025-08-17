#!/bin/bash
set -e

# Process log files to extract errors and warnings
ls -1 logs/*.log \
| while read -r file; do
  if ! [[ -f "${file}" ]]; then
    continue
  fi
  echo "Processing ${file}"
  grep -i "error\|warning" "${file}" \
  | while read -r line; do
    timestamp=$(echo "${line}" | cut -d' ' -f1)
    level=$(echo "${line}" | cut -d' ' -f2)
    echo "${timestamp},$level" >> results.csv
  done
done
