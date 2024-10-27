#!/bin/bash

DIRECTORY="./db/sqlc"

for FILE in "$DIRECTORY"/*; do
    if [ -f "$FILE" ]; then
        awk '/package db/ {found=1} found {print}' "$FILE" > temp && mv temp "$FILE"
        echo "Processed $FILE"
    fi
done

echo "Processing complete."
