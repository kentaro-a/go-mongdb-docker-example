mongoimport --authenticationDatabase admin -u root -p mongo --db test --collection users --file /docker-entrypoint-initdb.d/users.json --jsonArray

