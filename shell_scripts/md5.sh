#!/bin/bash

# Prompt user for MariaDB username
read -p "Enter MariaDB username: " db_username

# Prompt user for MariaDB password (without showing the input)
read -s -p "Enter MariaDB password: " db_password
echo

# Prompt user for username
read -p "Enter the username: " username

# Prompt user for password (without showing the input)
read -s -p "Enter the password: " password
echo

# Calculate the MD5 hash of the password
md5hash=$(echo -n "$password" | md5sum | awk '{print $1}')

# Construct the SQL query
sql_query="USE TGSCA; INSERT INTO Users (Username, Password, IsAdmin, PlainPassword) VALUES ('$username','$md5hash',true,'$password');"

# Run the SQL query using the mysql command-line tool
echo "$sql_query" | mysql -u "$db_username" -p"$db_password"

echo "SQL query executed successfully!"