# aggreGATOR



notes:  
How to get into gator psql?  
$`sudo -u postgres psql`  
`\c gator`  
`SELECT * FROM users;`

How to migrate down and back up with goose?  
-goose postgres <connection_string> up-  
cd into the sql/schema directory and run:  
`goose postgres "postgres://postgres:postgres@localhost:5432/gator" down`   
`goose postgres "postgres://postgres:postgres@localhost:5432/gator" up`   

How to register a user?  
`go run . register lane`

How to login as a registered user?  
`go run . login lane`
