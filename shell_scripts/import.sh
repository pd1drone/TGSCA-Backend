#!/bin/bash

# Prompt the user for their MariaDB username
read -p "Enter your MariaDB username: " username

# Prompt the user for their MariaDB password (the -s flag hides the input)
read -s -p "Enter your MariaDB password: " password
echo

# Run the database.sql script using the entered username and password
mysql -u "$username" -p"$password" < EnrollementSystem.sql