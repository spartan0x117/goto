# goto

The general idea is to allow for quickly visiting urls based on easily-remembered and discovered labels, either in the browser or via the CLI.
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

## Setup using links in the browser
- First, you must modify `/etc/hosts` to include `goto` and `www.goto` (or whatever you would like to type into the browser search bar as a prefix for the search, e.g. `goto/github` or `go/github`) as a local resolution. A minimal example:
```
127.0.0.1   localhost goto www.goto
::1         localhost goto www.goto
```
- Next, in order to be able to bind to port 80, you may need to make some changes (in testing, on Mac OS, no changes were needed. On Ubuntu 22.04, changes were needed). There are a few options:
  - The simplest is to use `setcap` (you can read `man capabilities` and `man setcap` to learn more) to allow `goto server` to bind to priveleged ports. This can be done by running `sudo setcap 'CAP_NET_BIND_SERVICE=+ep' /path/to/goto/binary` (the path for the binary is likely `/usr/local/bin/goto`). You will unfortunately have to do this any time that the file is changed (such as installing a new version).
  - Alternatively, and this would be an exercise for the user, you can run nginx locally as root as a proxy.

## API Endpoints / Using links in the browser
You can access the following in your browser's search bar to use `goto` most efficiently!

_NOTE: The following assumes the name you gave in `/etc/hosts` is `goto`, but you can name it whatever you want._

_NOTE: What you type into the url bar must be recognized as a valid url for the browser to perform a DNS lookup (for Firefox and Chrome, at least). It may be possible to get around this with some browser settings (like `browser.fixup.dns_first_for_single_words` for Firefox, though this comes with downsides), but the simplest way is to either append a `/` to every query (i.e. `goto/example/` instead of `goto/example`) or to set something in `/etc/hosts` that is a valid url (i.e. `go.to/example`)_


- `goto/<label>/`: Redirects to the url for the particular label, or returns `404 Not Found` if it does not exist.
- `goto/add/`: Opens a page that allows creating a new link. 
  - Note that if the configured storage is `git`, adding may take a few seconds as it is performing both a `git pull` and `git push`.
- `goto/find[/<label>]/`: With no label in the path (i.e. `goto/find`), lists all labels. If a label is present in the path (i.e. `goto/find/example`), returns the url for the label without redirecting.
- `goto/remove/<label>/`: Removes the link for `label`, or returns `404 Not Found` if it does not exist. 
  - Note that if the configured storage is `git`, removing may take a few seconds as it is performing both a `git pull` and `git push`.
- `goto/sync/`: Syncs with configured storage.

## Commands

- `add <label> <url>`: adds a new entry to the link store with the specified label and url. If an entry with the same label already exists,
  a prompt appears with a choice whether to update the entry or not.
  - NOTE: for a `git_config` add will always pull the configured repo before running.
- `alias`
  - `alias add <alias> <label>`: adds a new local alias for a label. The alias file is found at `~/.config/goto/aliases.json`.
  - `alias remove <alias>`: removes a local alias.
- `find [<label>]`: displays the url for the specified label, if any. If no label is specified, will display *all* labels in the store. 
  Combine this with `grep` to find a label based on a regex!
- `init`: explicitly create initial config files for goto. Will do nothing if they already exist.
- `open <label>`: opens the url for the label in a browser. It is possible to supply a path to be used with the url.
  - For example: if the entry `jira:https://jira.com` existed, one could execute `goto jira/ops/123` and this would result in the url
    `https://jira.com/ops/123` being opened in the browser.
  - NOTE: Running `goto <label>` is equivalent to running `goto open <label>`. This does mean that a label cannot have the same name as any of `goto`'s subcommands.
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

Create config directory and files for goto:

```bash
goto init
```

Files will be created in `~/.config/goto/`.


## Ideas for future development

*There is no official roadmap, but these are some ideas for future improvements.*

- [ ] Auxiliary local/private store for the git-based store. This would allow for adding links that are not pushed to the git remote and
  are only for the current user.
- [X] Local aliases. Similar to the above, it would allow the user to create aliases for labels that are not pushed to a remote. This would prevent polluting the shared repo with duplicate urls. What is short/convenient for one person may not be for another
  (e.g. say there is an existing `github:https://github.com` entry. Someone who uses it a lot may want to have 
  `gh:https://github.com` for quick access).
- [X] Initialization helpers (creating `~/.config/goto/` along with necessary files, setting up `links.json` in an empty repo).
- [ ] Label/command tab-completion.
- [X] Server mode. Would expose a locally running API that could be used by a browser directly.
