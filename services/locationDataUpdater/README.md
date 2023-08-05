# Location Data Updater.

## Packages

GORM for ORM in golang
* go get gorm.io/gorm
* go get gorm.io/driver/mysql

## WTF?

This service is responsable for updating filter database base on new spots that where created in other databases.

![now-Filter service drawio (1)](https://user-images.githubusercontent.com/21164304/196567236-50c457f3-b29d-4122-b708-476964d3b6fc.png)


## Useful links
* TBD


## Local test
docker run --name mysql -p 3306:3306 -p 33006:33006 -e MYSQL_ROOT_PASSWORD=admin  -e MYSQL_PASSWORD=admin -e MYSQL_USER=admin -e MYSQL_DATABASE=pululapp -d mysql