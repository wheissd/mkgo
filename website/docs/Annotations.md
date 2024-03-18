---
sidebar_position: 20
---

## Available annotations:

Every annotation func accepts opt argumnets.
Currently only modes option is available - it gives availability to configure in which modes(Mode is field in config) you enable or disable this option.
```go {5}
// Edges of the Cat.
func (Cat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("kittens", Kitten.Type).Annotations(
			annotations.Edge().EnableRead(annotations.Modes{"default"}),
		),
		edge.From("breed", Breed.Type).
			Required().
			Unique().
			Field("breed_id").
			Ref("cats"),
	}
}
```
### Field

#### SetPublic
Sets field public
```go {4}
func (Cat) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
		Annotations(annotations.Field().SetPublic(),
    }
}
```
#### SetPrivate
Sets field private
Sets field public
```go {4}
func (Cat) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
		Annotations(annotations.Field().EnableFilter(),
    }
}
```
#### EnableFilter
Enables filtering by this field
```go {4}
func (Cat) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
		Annotations(annotations.Field().EnableFilter(),
    }
}
```

### Entity
#### EnableCreate
Enable create operation for entity
```go {3}
func (Cat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		annotations.Entity().EnableCreate(annotations.Modes{"api"}),
	}
}
```
#### DisableCreate
Disable create operation for entity
```go {3}
func (Cat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		annotations.Entity().DisableCreate(annotations.Modes{"api"}),
	}
}
```
#### EnableUpdate
Enable update operation for entity
```go {3}
func (Cat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		annotations.Entity().EnableUpdate(annotations.Modes{"api"}),
	}
}
```
#### DisableUpdate
Disable update operation for entity
```go {3}
func (Cat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		annotations.Entity().DisableUpdate(annotations.Modes{"api"}),
	}
}
```
#### EnableDelete
Enable delete operation for entity
```go {3}
func (Cat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		annotations.Entity().EnableDelete(annotations.Modes{"api"}),
	}
}
```
#### DisableDelete
Disables delete operation for entity
```go {3}
func (Cat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		annotations.Entity().DisableDelete(annotations.Modes{"api"}),
	}
}
```
#### EnableReadOne
Enable read one operation for entity
```go {3}
func (Cat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		annotations.Entity().EnableReadOne(annotations.Modes{"api"}),
	}
}
```
#### DisableReadOne
Disable read one operation for entity
```go {3}
func (Cat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		annotations.Entity().DisableReadOne(annotations.Modes{"api"}),
	}
}
```
#### EnableReadMany
Enable read many operation for entity
```go {3}
func (Cat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		annotations.Entity().EnableReadMany(annotations.Modes{"api"}),
	}
}
```
#### DisableReadMany
Disable read many operation for entity
```go {3}
func (Cat) Annotations() []schema.Annotation {
	return []schema.Annotation{
		annotations.Entity().DisableReadMany(annotations.Modes{"api"}),
	}
}
```

### Edge
#### EnableRead
Enable read operation for edge
```go {5}
// Edges of the Cat.
func (Cat) Edges() []ent.Edge {
return []ent.Edge{
        edge.To("kittens", Kitten.Type).Annotations(
            annotations.Edge().EnableRead(annotations.Modes{"default"}),
        ),
    }
}

```
#### DisableRead
Disable read operation for edge
```go {5}
// Edges of the Cat.
func (Cat) Edges() []ent.Edge {
return []ent.Edge{
        edge.To("kittens", Kitten.Type).Annotations(
            annotations.Edge().DisableRead(annotations.Modes{"default"}),
        ),
    }
}

```
#### EnableWrite
Enable write operation for edge
```go {5}
// Edges of the Cat.
func (Cat) Edges() []ent.Edge {
return []ent.Edge{
        edge.To("kittens", Kitten.Type).Annotations(
            annotations.Edge().EnableWrite(annotations.Modes{"default"}),
        ),
    }
}

```
#### DisableWrite
Disable write operation for edge
```go {5}
// Edges of the Cat.
func (Cat) Edges() []ent.Edge {
return []ent.Edge{
        edge.To("kittens", Kitten.Type).Annotations(
            annotations.Edge().DisableWrite(annotations.Modes{"default"}),
        ),
    }
}

```
