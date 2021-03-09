# build docker image
docker-build:
	sudo docker build -t hello .

# run docker image
docker-run:
	sudo docker run -d -p 9000:9000 hello

# start traefik
docker-up-proxy:
	sudo docker-compose up -d reverse-proxy

# make instance for hello service
docker-up-hello:
	sudo docker-compose up -d --scale hello=3
