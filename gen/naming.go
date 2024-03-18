package gen

func schemaEntityNameRead(n string) string {
	return n + "Read"
}

func schemaEntityNameAsIs(n string) string {
	return n
}

func schemaEntityNameCreate(n string) string {
	return n + "Create"
}

func schemaEntityNameUpdate(n string) string {
	return n + "Update"
}
