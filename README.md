# TGSCA-Backend

To Setup:

1. Install a debian virtual machine in virtualbox
2. After installing update and install required packages but change to root user first.
change to root user command:
```
su
```
> then enter the password of the user root
Update command:
```
sudo apt update
```
Install required packges command:
```
apt install wget git ssh make
```
> then you need to configure the ssh to enable to login using root remotely.
run this command to enable ssh and enable root login remotely
```
sudo sed -i 's/#PermitRootLogin prohibit-password/PermitRootLogin yes/' /etc/ssh/sshd_config
```
then next you need to restart ssh
```
systemctl restart ssh
```
3. go to the /root directory git clone the TGSCA-Backend
```
git clone https://github.com/pd1drone/TGSCA-Backend
```
4. Go to the directory TGSCA-Backend
```
cd TGSCA-Backend/
```
5. Make the shell scripts executable
```
chmod +x shell_scripts/create_tgsca_service_file.sh
```
```
chmod +x shell_scripts/import.sh
```
```
chmod +x shell_scripts/install.sh
```
```
chmod +x shell_scripts/md5.sh
```
6. run the install.sh script
```
shell_scripts/install.sh
```
then next run this command to be able to run go
```
source /etc/profile.d/go.sh
```
7. run the import.sh script
```
shell_scripts/import.sh
```
8. run the make command to tidy the packages and create a golang executable file
```
make build
```
9. run the create_tgsca_service_file.sh
```
shell_scripts/create_tgsca_service_file.sh
```

## TO Check if the Backend is now running run the command
```
systemctl status tgsca_backend
```

## Also access this link to check if you have access to the backend api
```
http://<ipaddress_of_debian_vm>:8082/check
```