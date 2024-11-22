---
sidebar_position: 100
---

# FAQ

### Crud operations are not generated at all. What's the problem?
If you use options.OpenapiSchema and for example reset spec.Paths - then all generated paths will be missed, don't delete paths if you don't want to.
