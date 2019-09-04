# git-ls

Explore github content given a valid token.

## Usage

The token is read from the `GITHUB_TOKEN` environment variable. I chose this so I don't end up
putting tokens in HISTFILEs. Your options are to source it from a file with `source .env`, export
it in your shell session, or go ahead and prepend it like so: `GITHUB_TOKEN=xyz git-ls user`.

Run `git-ls` with no arguments to see what data can be listed.

## Examples

Below are some useful filters to use with `jq`, the CLI json parser.

```bash
# Just private gists
❯❯ git-ls gists | jq '.[] | select(.Private == true)'

# Just repo names
❯❯ git-ls repos | jq -r '.[].Name'

# Repos that don't belong to the GITHUB_TOKEN owner
❯❯ git-ls repos | jq '.[] | select(.Owner != "audibleblink")'
# or
❯❯ git-ls collabs
```

## Properties

I tried to select the most useful properties returned by the GitHub API to keep noise low.
These are the properties that you can filter on with `jq`

### Gists

* Owner
* Description
* GitPullUrl
* Files
* Private

### Repos

* Name
* Description
* URL
* Owner
* Organization
* StargazersConut
* Private


## Accessing Data
If you have the API token that got you this far, 
it can be placed into an HTTP clone request.

```bash
git clone https://${GITHUB_TOKEN}@github.com/someOrg/someRepo
```

If you just want all the secret repos the token has access to

```bash
git-ls plunder
```
![](https://i.imgur.com/lcn6Wop.png)
![](https://i.imgur.com/s587JPU.png)
