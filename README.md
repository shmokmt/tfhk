# tfhk

The utility tool to remove blocks for refactoring such as moved block.

Supports deletion of the following blocks.

- moved block
- import block
- removed block

# Usage

```
Usage: tfhk [-recursive] [target]
  -recursive
        Also process files in subdirectories. By default, only the given directroy (or current directroy) is processed.
```

# References

- https://developer.hashicorp.com/terraform/language/modules/develop/refactoring#removing-moved-blocks
