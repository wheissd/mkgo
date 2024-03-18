---
sidebar_position: 1
---

# Relations

To enable read with relations included use **annotations.Edge().EnableRead()**

```go title='your_model_schema.go' {5}
// Edges of the Cat.
func (Cat) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("kittens", Kitten.Type).Annotations(
			annotations.Edge().EnableRead(),
		),
	}
}
```
