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

You can also make a bot create a pull request using GitHub Actions.
An example workflow is shown below.

> [!TIP]
> If you want to trigger other actions, please use GitHub Apps or PAT.
> see also https://github.com/orgs/community/discussions/65321

```yaml
name: tfhk-pr

on:
  workflow_dispatch:
  schedule:
   - cron: '0 1 * * 1'

permissions:
  contents: write
  pull-requests: write

jobs:
  create-pull-request:
    runs-on: ubuntu-latest
    steps:
    - name: Checkout repository
      uses: actions/checkout@v4

    - uses: actions/setup-go@v5
      with:
        go-version: 'stable'

    - uses: actions/create-github-app-token@v1
      id: app-token
      with:
        app-id: ${{ secrets.TFHK_APP_ID }}
        private-key: ${{ secrets.TFHK_PRIVATE_KEY }}

    - name: Install tfhk
      run: go install github.com/shmokmt/tfhk/cmd/tfhk@v0.2.0

    - name: Run tfhk
      run: tfhk -recursive .

    - name: Check changes
      id: diff-check
      run: git diff --exit-code || echo "changes_detected=true" >> $GITHUB_OUTPUT

    - name: Commit changes
      if: steps.diff-check.outputs.changes_detected == 'true'
      run: |
        echo steps.diff-check.outputs.changes_detected: ${{ steps.diff-check.outputs.changes_detected }}
        branch_name=tfhk_$(date +"%Y%m%d%H%M")
        git switch -c ${branch_name}
        git config --global user.email "41898282+github-actions[bot]@users.noreply.github.com"
        git config --global user.name "github-actions[bot]"
        git add .
        git diff --cached --exit-code || (git commit -m "auto-remove blocks by tfhk" && git push origin ${branch_name})
        gh pr create --base main --head ${branch_name} --title "auto-remove blocks by tfhk" --body ""
      env:
        GH_TOKEN: ${{ steps.app-token.outputs.token }}
```

# References

- https://developer.hashicorp.com/terraform/language/modules/develop/refactoring#removing-moved-blocks
