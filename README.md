# hackernews

Using grapQL and gola

## Init graQL

go get github.com/99designs/gqlgen
go run github.com/99designs/gqlgen init
go run ./server.go

## Init mysql

go get -u github.com/go-chi/chi/v5
go get -u github.com/go-sql-driver/mysql
go get github.com/golang-migrate/migrate/v4/cmd/migrate
go build -tags 'mysql' -ldflags="-X main.Version=1.0.0" -o $GOPATH/bin/migrate github.com/golang-migrate/migrate/v4/cmd/migrate/
migrate create -ext sql -dir mysql -seq create_users_table
migrate create -ext sql -dir mysql -seq create_links_table
migrate -database mysql://root:example@/hackernews -path ./mysql up

## Get data

query {
	links{
    title
    address,
    user{
      name
    }
  }
}

curl 'http://localhost:8080/query' \
  -H 'Accept-Language: en-US,en;q=0.9' \
  -H 'Connection: keep-alive' \
  -H 'Origin: http://localhost:8080' \
  -H 'Referer: http://localhost:8080/' \
  -H 'Sec-Fetch-Dest: empty' \
  -H 'Sec-Fetch-Mode: cors' \
  -H 'Sec-Fetch-Site: same-origin' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36' \
  -H 'accept: application/json, multipart/mixed' \
  -H 'content-type: application/json' \
  -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="102", "Google Chrome";v="102"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  --data-raw '{"query":"query {\n\tlinks{\n    title\n    address,\n    user{\n      name\n    }\n  }\n}","variables":null}' \
  --compressed

## Mution

mutation {
  createLink(input: {title: "new link", address:"http://address.org"}){
    title,
    user{
      name
    }
    address
  }
}

curl 'http://localhost:8080/query' \
  -H 'Accept-Language: en-US,en;q=0.9' \
  -H 'Connection: keep-alive' \
  -H 'Origin: http://localhost:8080' \
  -H 'Referer: http://localhost:8080/' \
  -H 'Sec-Fetch-Dest: empty' \
  -H 'Sec-Fetch-Mode: cors' \
  -H 'Sec-Fetch-Site: same-origin' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36' \
  -H 'accept: application/json, multipart/mixed' \
  -H 'content-type: application/json' \
  -H 'sec-ch-ua: " Not A;Brand";v="99", "Chromium";v="102", "Google Chrome";v="102"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  --data-raw '{"query":"mutation {\n  createLink(input: {title: \"new link\", address:\"http://address.org\"}){\n    title,\n    user{\n      name\n    }\n    address\n  }\n}","variables":null}' \
  --compressed

## Doc

[source](https://github.com/howtographql/graphql-golang/blob/master/graph/schema.graphqls)
[doc](https://www.howtographql.com/graphql-go/8-logged-in-user-object/)

migrate -database mysql://root:example@/hackernews -path ./mysql up

## Access to mysql

mysql -u root -h localhost -P 3306 hackernews -p
