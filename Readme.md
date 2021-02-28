# HOW TO INSTALL

##### 1. Clone the repository

```bash
git clone https://github.com/feedlyy/Stockbit.git
```

##### 2. Install Prerequisites
This Project use Go 1.14, please download the same version
or the newer version of Go from https://golang.org/dl/

- mysql driver
```
https://github.com/go-sql-driver/mysql
```

This project use mysql as database. if you didn't have mysql
in your machine, you can run mysql with this command:
```
make mysql
```
this command will run database with these credentials:
```
user = root
password = secret
host = localhost
port = 3306
```

#### 3. Answer
Answer for this Coding Test online for number 1, 3, and 4
are in Answer1-3-4 folder, please take a look.

Answer for Microservice are all files in this folder except
Answer1-3-4 folder, the routes are in main.go .
