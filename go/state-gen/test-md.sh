#!/bin/sh

# Run test files from the test directory
if [ -d "../../test" ]; then
    for file in ../../test/*; do
        if [ -f "$file" ]; then
            echo "\nProcessing file: $file"
            echo "-------------------"
            cat "$file" | go run state-gen.go
        fi
    done
else
    echo "Error: test directory not found"
    exit 1
fi
