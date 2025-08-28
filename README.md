# aggreGATOR



notes:  
### How to generate the Go code?  
Write a query to create a user. Inside the sql/queries directory, create a file called users.sql.

Run $`sqlc generate` from the root of your project.  
It should create a new package of go code in `internal/database`.  
files: db.go, models.go, users.sql.go  

### How to get into gator psql?  
$`sudo -u postgres psql`  
`\c gator`  
`SELECT * FROM users;`

### How to migrate down and back up with goose?  
-goose postgres <connection_string> up-  
cd into the sql/schema directory and run:  
`goose postgres "postgres://postgres:postgres@localhost:5432/gator" down`   
`goose postgres "postgres://postgres:postgres@localhost:5432/gator" up`  
useful after making a new table, resting the db (there's a reset command now)

### How to register a user?  
`go run . register lane`

### How to login as a registered user?  
`go run . login lane`

### How to migrations
made a new table in sql\schema\<file>

