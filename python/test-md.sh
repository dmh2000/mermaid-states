#!/bin/bash

for file in ../test/*.md; do
    python3 state-gen.py < "$file"
done