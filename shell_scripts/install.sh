#!/bin/bash

# Install MariaDB
sudo apt update
sudo apt install -y mariadb-server

# Start and enable MariaDB daemon
sudo systemctl start mariadb
sudo systemctl enable mariadb

# Secure the installation
sudo mysql_secure_installation <<EOF

y
admin123
admin123
y
y
y
y
EOF

echo "MariaDB installation and setup completed!"
