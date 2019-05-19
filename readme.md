# git-ls

Explore github content given a valid token.

## Usage

The token is read from the `GITHUB_TOKEN` environment variable. I chose this so I don't end up
putting tokens in HISTFILEs. So your options are to source it from a file with `source .env` or
go ahead and prepend it like so: `GITHUB_TOKEN=xyp git-ls`.


```bash
repos
acls
orgs
user
```
