#!/bin/bash

# Create the file tgsca_backend.service with the specified content
cat << EOF > /etc/systemd/system/tgsca_backend.service
[Unit]
Description=TGSCA backend http server
After=mariadb.service

[Service]
Restart=always
User=root
WorkingDirectory=/root/TGSCA-Backend/
ExecStart=/root/TGSCA-Backend/cmd/tgsca
Requires=mariadb.service

[Install]
WantedBy=multi-user.target
EOF

# Reload the systemd daemon
sudo systemctl daemon-reload

# Enable the tgsca_backend.service file so it starts on boot
sudo systemctl enable tgsca_backend.service

# Reload the systemd daemon again to pick up the changes from enabling the service
sudo systemctl daemon-reload

# Start the tgsca_backend.service
sudo systemctl start tgsca_backend.service
