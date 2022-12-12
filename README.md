# goto
A CLI alternative to go/ links. The general idea is to allow for quickly visiting urls based on easily-remembered and discovered labels. 
These `label->url` pairs can be collaboratively updated, making it more powerful the more people in your group/org/company that use it. 

# Configuration
~/.config/goto/config.yaml
```
type: "git|json"
git_config:
  local_path: /path/to/repo/on/local/machine
  auto_sync: true|false                             # If true, goto will pull on every operation instead of only on add/remove
json_config:
  path: /path/to/local/file
```

For a git config, you need to have checked out a git repository containing a file named `links.json`

# Commands
- `add <label> <url>`: adds a new entry to the link store with the specified label and url. If an entry with the same label already exists,
  a prompt appears with a choice whether to update the entry or not.
    - NOTE: for a `git_config` add will always pull the configured repo before running.
- `alias`
  - `alias add <alias> <label>`: adds a new local alias for a label. The alias file is found at `~/.config/goto/aliases.json`.
  - `alias remove <alias>`: removes a local alias.
- `find [<label>]`: displays the url for the specified label, if any. If no label is specified, will display *all* labels in the store. 
  Combine this with `grep` to find a label based on a regex!
- `open <label>`: opens the url for the label in a browser. It is possible to supply a path to be used with the url.
  - For example: if the entry `jira:https://jira.com` existed, one could execute `goto jira/ops/123` and this would result in the url
    `https://jira.com/ops/123` being opened in the browser.
  - NOTE: Running `goto <label>` is equivalent to running `goto open <label>`. This does mean that a label cannot have the same name as any 
    of `goto`'s subcommands.
- `remove <label>`: removes the entry for the specified label.
  - NOTE: for a `git_config` remove will always pull the configured repo before running.
- `sync`: syncs with the configured remote. For a `json_config` this is a no-op. For a `git_config`, this pulls from the configured remote.

# Build/develop/setup locally
At the root of the repository, run `make build` to build a binary to `out/goto`. Move this binary to a directory (like `/usr/local/bin`) in your 
`PATH` to use it easily from the command line.

Run `mkdir -p ~/.config/goto/`

Add your config at `~/.config/goto/config.yaml`

# Ideas for future development
_There is no official roadmap, but these are some ideas for future improvements._
- [ ] Auxiliary local/private store for the git-based store. This would allow for adding links that are not pushed to the git remote and
  are only for the current user.
- [X] Local aliases. Similar to the above, it would allow the user to create aliases for labels that are not pushed to a remote. This 
  would prevent polluting the shared repo with duplicate urls. What is short/convenient for one person may not be for another
  (e.g. say there is an existing `github:https://github.com` entry. Someone who uses it a lot may want to have 
  `gh:https://github.com` for quick access).
- [ ] Initialization helpers (creating `~/.config/goto/`, setting up `links.json` in an empty repo).
- [ ] Label/command tab-completion.
