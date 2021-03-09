# build docker image
sudo docker build -t hello .

# run docker image
sudo docker run -d -p 9000:9000 hello

# start traefik
sudo docker-compose up -d reverse-proxy

# make instance for hello service
sudo docker-compose up -d --scale hello=3