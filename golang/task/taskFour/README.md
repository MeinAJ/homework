go mod init blog && go mod tidy

go mod init blog
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

docker network create mynetwork

docker build -t myapp:1.0 .

docker run -d --name myapp --network mynetwork -p 8080:8080 myapp:1.0

docker run -d --name nacos --network mynetwork -p 8848:8848 -e MODE=standalone nacos/nacos-server:v2.2.0

docker run -d --name redis-single-node --network mynetwork -p 6379:6379 redis:7.4.2

docker run -d --name mysql-container --network mynetwork -p 3306:3306 -v /Users/ja/tmp/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=root mysql:latest

docker run -d -p 2379:2379 -p 2380:2380 --network mynetwork -e ETCDCTL_API=3 --mount type=bind,source=/d/tmp/etcd/single,destination=/tmp/etcd --name etcd-single gcr.io/etcd-development/etcd:v3.4.25 /usr/local/bin/etcd --config-file /d/tmp/etcd/etcd.yml

# install nvm
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.5/install.sh | bash

# install node
nvm install 19.0.0

# use node
nvm use 19.0.0