# aggreGATOR

A command-line RSS/Atom feed aggregator written in Go with a PostgreSQL backend.  
You can register users, follow feeds, and browse new items — all from the terminal.

---

## Requirements

- **Go 1.21+** (for building/installing)
- **PostgreSQL** (running locally or accessible over network)

---

## Installation

Build and install the `gator` CLI:

```bash
go install github.com/CodyMcCarty/aggreGATOR@latest
```

--- 

## Configuration

gator expects a configuration file that tells it how to connect to your database.  
1. Create a config file at ~/.gatorconfig.json:
```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user": ""
}
```
1. Run database migrations (uses goose)
```bash
`goose postgres "postgres://postgres:postgres@localhost:5432/gator" up`  
```

--- 

## Usage

During development you can use:
```bash
go run . <command> [args]
```

But for production and normal use you should just run the compiled binary:
```bash
gator <command> [args]
```

### Common Commands
Register a user
```bash
gator register <username>


Login as a user
```bash
gator login <username>
```

Add a feed
```bash
gator addfeed <url>
```

Follow a feed
```bash
gator follow <url>
```

Browse recent items
```bash
gator browse
```

Start the aggregator loop
```bash
gator agg 30s
```
Runs the poller every 30 seconds. Duration accepts units like ms, s, m, h.

--- 

## Notes
- Go produces statically compiled binaries. Once you’ve built or installed gator, you don’t need Go itself to run the program. 
- go run . is best for development. End-users should use the gator binary.

--- 

## personal notes and reminders:  
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



