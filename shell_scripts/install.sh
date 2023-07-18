#!/bin/bash

# Install MariaDB
sudo apt update
sudo apt install -y mariadb-server

# Start and enable MariaDB daemon
sudo systemctl start mariadb
sudo systemctl enable mariadb

# Secure the installation
sudo mysql_secure_installation <<EOF

n
y
admin123
admin123
y
n
n
y
EOF

echo "MariaDB installation and setup completed!"


echo "Installing Golang...."
cd /root
wget https://go.dev/dl/go1.20.6.linux-amd64.tar.gz
sudo tar -zxvf go1.20.6.linux-amd64.tar.gz -C /usr/local/
echo "export PATH=/usr/local/go/bin:${PATH}" | sudo tee /etc/profile.d/go.sh
source /etc/profile.d/go.sh
cd /root/TGSCA-Backend/
go version

echo "Finished Installing Golang"

