# GO GIT GITEA

The swiss army knife of Gitea. Useful functions for automating Gitea.

## Supported Functions

- Create mirrors of your existing
  - [ ] Github repositories
  - [ ] Gitlab repositories

- More TBD

## Getting Started

```
go build -o go-git-gitea main.go
```

## Special ENV Vars

| Variable | Default | Description |
| :------- | :------ | :---------- |
| `GT_URL` | https://gitea.com | The base URL of your Gitea instance
| `GT_API_TOKEN` | None | Your Gitea API Token. Found under 
| `GL_URL` | https://gitlab.com | The base URL of your Gitlab instance
| `GL_API_TOKEN` | None | Your Gitlab API Token. Found under Preferences >> Access Tokens
| `GH_URL` | https://github.com | The base URL of your Github instance
| `GH_API_TOKEN` | None | Your Github API Token. Found under Settings >> Developer Settings >> Personal Access Tokens

These variables can also be set via the CLI if necessary. CLI input will override ENV vars.
