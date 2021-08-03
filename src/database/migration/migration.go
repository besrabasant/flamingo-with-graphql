package migration

type MigrationUp interface {
	up()
}

type MigrationDown interface {
	down()
}
