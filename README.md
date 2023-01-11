# goto

A CLI alternative to go/ links. The general idea is to allow for quickly visiting urls based on easily-remembered and discovered labels.
These `label->url` pairs can be collaboratively updated, making it more powerful the more people in your group/org/company that use it.

## Configuration

~/.config/goto/config.yaml

```yaml
type: "git|json"
git_config:
  local_path: /path/to/repo/on/local/machine
  auto_sync: true|false                             # If true, goto will pull on every operation instead of only on add/remove
json_config:
  path: /path/to/local/file
```

For a git config, you need to have checked out a git repository containing a file named `links.json`

### Using links in the browser
- First, you must modify `/etc/hosts` to include `goto` and `www.goto` (or whatever you would like to type into the browser search bar as a prefix for the search, e.g. `goto/github` or `go/github`) as a local resolution. A minimal example:
```
127.0.0.1   localhost goto www.goto
::1         localhost goto www.goto
```
- Next, for Firefox you must set `browser.fixup.dns_first_for_single_words` to `true`. 
  - You can do this by opening a Firefox window and typing `about:config` and hitting return. Then, in the search bar, type `browser.fixup.dns_first_for_single_words` and click the button on the right-hand side.
- For Google Chrome no settings should need updating. However, you can test that the resolution is being applied by going to `chrome://net-internals/#dns` and doing a DNS lookup for `goto` (or whatever you entered for the prefix) in the `DNS` section.

## Commands

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
- `server`: Runs a server locally on port 80. This is intended to be used as a proxy for use in the browser. 
- `sync`: syncs with the configured remote. For a `json_config` this is a no-op. For a `git_config`, this pulls from the configured remote.

## Build/develop/setup locally

At the root of the repository, run

```bash
make build
```

to build a binary to `out/goto`.  To install the binary, run:

```bash
sudo make install
```

By default, it will be installed in `/usr/local/bin`. You can install it privately by overriding the destination directory:

```bash
make install DEST=<path>
 ```

 where `<path>` is a directory in `$PATH` and is writable by you.

Create a config directory for goto:

```bash
mkdir -p ~/.config/goto/
```

Add your config at `~/.config/goto/config.yaml`

Create an initial empty `{}` map in `~/.config/goto/aliases.json` for storing aliases

## Ideas for future development

*There is no official roadmap, but these are some ideas for future improvements.*

- [ ] Auxiliary local/private store for the git-based store. This would allow for adding links that are not pushed to the git remote and
  are only for the current user.
- [X] Local aliases. Similar to the above, it would allow the user to create aliases for labels that are not pushed to a remote. This would prevent polluting the shared repo with duplicate urls. What is short/convenient for one person may not be for another
  (e.g. say there is an existing `github:https://github.com` entry. Someone who uses it a lot may want to have 
  `gh:https://github.com` for quick access).
- [ ] Initialization helpers (creating `~/.config/goto/` along with necessary files, setting up `links.json` in an empty repo).
- [ ] Label/command tab-completion.
- [X] Server mode. Would expose a locally running API that could be used by a browser directly.
