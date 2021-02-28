.PHONY: mysql migrate

mysql:
	docker run --name mysql -p 3306:3306 -e NAME=root -e PASSWORD=secret -e MYSQL_ROOT_PASSWORD=secret -d mysql:5.7.33