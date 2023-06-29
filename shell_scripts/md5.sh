#!/bin/bash

# Run the command and store the output in a variable
output=$(echo -n 'admin123' | md5sum)

# Extract the hash value from the output
adminPassword=$(echo "$output" | awk '{print $1}')

# Print the extracted hash value
echo $adminPassword