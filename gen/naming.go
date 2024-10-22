package gen

func schemaEntityNameRead(n string) string {
	return "Read" + n
}

func schemaEntityNameAsIs(n string) string {
	return n
}

func schemaEntityNameCreate(n string) string {
	return "Create" + n
}

func schemaEntityNameUpdate(n string) string {
	return "Update" + n
}
