db:
	docker run --name mysql --rm -p 3306:3306 -e MYSQL_ROOT_PASSWORD=p4ssw0rd mysql:latest -a

clean:
	docker stop mysql && docker rmi mysql && docker rm -f $(docker ps -a -q)

containers:
	docker container ls