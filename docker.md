**installing inside wsl**

https://docs.docker.com/engine/install/ubuntu/

https://linuxhandbook.com/system-has-not-been-booted-with-systemd/

service docker start


**Static linking**

https://www.blang.io/posts/2015-04_golang-alpine-build-golang-binaries-for-alpine-linux/

in .profile

EXPORT CGO_ENABLED=0


**check if app as links to libraries**

ldd \<executable filename\>


**build image**

docker build -t \<tag\> .


**run**

docker run -it --rm -p 5051:5050 \<tag\>



**push**

docker login \<url\>

docker tag \<local tag\>  \<url/tag\>

docker push \<url/tag\>

**need to add timezone for Alpine**

https://github.com/takecy/tz-sample

add to Dockerfile

RUN apk --no-cache add tzdata


**Give non-root user permission to run docker**

sudo usermod -aG docker ${USER}


**MySQL**

docker run -p 127.0.0.1:3306:3306 --name mariadb -e MYSQL_ROOT_PASSWORD=mypass -d mariadb

docker exec -it mariadb  /bin/bash

**Change system timezone which affects MySQL**

dpkg-reconfigure tzdata


**Remove image, build image, tag image, push image**

docker rmi web

docker build --tag web .

docker tag web <>.azurecr.io/web

docker push <>.azurecr.io/web
