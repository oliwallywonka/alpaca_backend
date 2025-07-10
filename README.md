# try JET
jet -dsn=postgresql://root:root@localhost:5432/alpaca?sslmode=disable -schema=public -ignore-tables=_prisma_migrations -path=./.gen

# just re-generate the Go client
go run github.com/steebchen/prisma-client-go generate
 
# sync the database with your schema for development
go run github.com/steebchen/prisma-client-go db push
 
# create a prisma schema from your existing database
go run github.com/steebchen/prisma-client-go db pull
 
# for production use, create a migration locally (THIS IS THE MOST IMPORTANT STEP)
go run github.com/steebchen/prisma-client-go migrate dev --name init --skip-generate
 
# sync your production database with your migrations
go run github.com/steebchen/prisma-client-go migrate deploy

goose up
goose create <migration name> sql

# JET GENERATE SQL BUILDER
jet -dsn=postgresql://root:root@localhost:5432/alpaca?sslmode=disable -schema=public -path=./.gen
jet -ignore-tables=_prisma_migrations -source=sqlite -dsn="file:dev.db" -path=./gen

# go run Ent init
go run -mod=mod entgo.io/ent/cmd/ent new

# Install entimport
go run -mod=mod ariga.io/entimport/cmd/entimport -h
# generate schemas
go run ariga.io/entimport/cmd/entimport -dsn "postgresql://root:root@localhost:5432/alpaca?sslmode=disable"