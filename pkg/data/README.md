## Database

We are going to use database engine called [CockroachDB(RDBM)](https://www.cockroachlabs.com/product/). It's a modern, Cloud-native, distributed SQL database. In this project, we use [Docker image for CockroachDB](https://hub.docker.com/r/cockroachdb/cockroach) and run it in container.

### User Schema

```
type User struct {
	gorm.Model
	Email			string	`gorm:"size:255;not null;unique"`
	Username		string	`gorm:"size:255;not null;"`
	PasswordHash	string	`gorm:"size:255;not null"`
	Fullname		string
	Role			int
}
```

By default, GORM pluralizes struct name to snake_cases as table name snake_case as column name and uses CreatedAt, UpdatedAt to track creating/updating time.

`Email` and `Username` should be unique in this service.

`gorm.Model`: Include ID, CreatedAt, UpdatedAt, DeletedAt
`PasswordHash`: Password should be hashed on client. In case of security breach, only hashed pass will be revealed.
`Role`: Used to define authentication levels and priority.(e.g. `0=standard` `1=admin`)

**Note**: Don't export whole list of users outside this package.
