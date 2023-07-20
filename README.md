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


#### TO CREATE A NEW ADMIN YOU NEED TO RUN THE shell_script/md5.sh
```
shell_scripts/md5.sh 
```
# You need to enter the mariadb username and password then you need to enter the username of the new admin and its password to be able to create this into the database.
OUTPUT:
```
root@debian-server:~/TGSCA-Backend# shell_scripts/md5.sh 
Enter MariaDB username: root
Enter MariaDB password: 
Enter the username: karl
Enter the password: 
SQL query executed successfully!
```

#### To check this in the database run the command

```
mysql -u root -p
```
> Then Enter the root password and run the SQL QUERY
```
USE TGSCA;
SELECT * FROM Users;
```

SAMPLE OUTPUT:
```
root@debian-server:~/TGSCA-Backend# mysql -u root -p
Enter password: 
Welcome to the MariaDB monitor.  Commands end with ; or \g.
Your MariaDB connection id is 7
Server version: 10.5.18-MariaDB-0+deb11u1 Debian 11

Copyright (c) 2000, 2018, Oracle, MariaDB Corporation Ab and others.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

MariaDB [(none)]> use TGSCA;
Reading table information for completion of table and column names
You can turn off this feature to get a quicker startup with -A

Database changed
MariaDB [TGSCA]> select * from Users;
+----+----------+----------------------------------+---------+---------------+
| ID | Username | Password                         | IsAdmin | PlainPassword |
+----+----------+----------------------------------+---------+---------------+
|  1 | admin    | 0192023a7bbd73250516f069df18b500 |       1 | NULL          |
|  8 | 123      | e16b2ab8d12314bf4efbd6203906ea6c |       0 | testpassword  |
| 23 | 20140001 | 1653a58dc41f5570a8806faf241104c6 |       0 | ZSHVwJYn      |
| 24 | 12345    | e563b704aeef61d0071be03cb8de74bb |       0 | XdBgh0mG      |
| 26 | karl     | e10adc3949ba59abbe56e057f20f883e |       1 | 123456        |
+----+----------+----------------------------------+---------+---------------+
5 rows in set (0.000 sec)

```

### You will see that a newly created admin in the Users Table