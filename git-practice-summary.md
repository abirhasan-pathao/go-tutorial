

---

## Practice Overview

| Branch | Purpose | Commands Practiced |
|--------|---------|-------------------|
| `main` | Main stable branch | `git checkout`, `git push` |
| `learn-git-changed-name` | Primary learning branch | `git branch -m <new-name>`[renamed the branch, pushed it upstream, removed the branch with old name from upstream], `git checkout`, squashed commits with `git rebase -i` |
| `learn-merge` | Practice merging | `git merge` [learned about fast-forward merge and merging after parent branch had moved forward] |
| `learn-rebase` | Practice rebasing | `git rebase`, `git rebase -i` |
| `learn-cherry-pick` | Practice cherry-picking | `git cherry-pick <commit-hash>`, `git cherry-pick <ref>`, `git cherry-pick <commit1>,<commit2>,<commit3>`, `git cherry-pick <commit1>^..<commitn>` |
| `learn-tag` | Practice tagging | `git tag`, `git tag -a`, `git push --tags` |
| `name-change` | Practice branch renaming | `git branch -m <old-name> <new-name>` |
| `pull-request` | Practice PR workflow | fetched from origin, rebased the branch, put out a pull request |
| `rebase-conflict` | Practice resolving rebase conflicts and conflict propagation | `git rebase`, resolved conflicts manually, `git rebase --continue` |

---

