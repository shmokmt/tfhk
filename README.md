# tfhk

[![Go](https://github.com/shmokmt/tfhk/actions/workflows/go.yml/badge.svg)](https://github.com/shmokmt/tfhk/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/shmokmt/tfhk.svg)](https://pkg.go.dev/github.com/shmokmt/tfhk)

The utility tool to remove blocks for refactoring such as moved blocks.

Supports deletion of the following blocks.

- moved block
- import block
- removed block

# Usage

```
go install github.com/shmokmt/tfhk/cmd/tfhk@latest
```

```
Usage: tfhk [-recursive] [target]
  -recursive
        Also process files in subdirectories. By default, only the given directroy (or current directroy) is processed.
```

# References

- https://developer.hashicorp.com/terraform/language/modules/develop/refactoring#removing-moved-blocks
