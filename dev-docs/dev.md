# Design
This document describes the current and future shape of this tool, it will
be a place for us, at Chef, to discuss about the design and usability, the
flow of commands that we expect user to follow plus, we will start the
documentation of this tool in an early stage.

## Help Documentation

### Main help
```bash
$ chef supermarket help
chef is a command-line tool that provides an interface between a local chef-repo and the Chef Infra Server. chef helps users to manage:
                        Nodes
                        Cookbooks and recipes
                        Roles, Environments, and Data Bags
                        Resources within various cloud environments
                        The installation of Chef Infra Client onto nodes
                        Searching of indexed data on the Chef Infra Server.

Usage:
  chef supermarket COMMAND ARTIFACT_TYPE ARTIFACT_NAME

Available Commands:
  help        Help about any command
  search      Search for cookbook on supermarket
  download    download cookbook from supermarket
  install     install cookbook on given path or /$HOMEDIR/.chef 

Flags:
  -m, --supermarket-site string  Supermarket URL
  
Use "chef supermarket [command] --help" for more information about a command.
$
```
### `search` sub-command help
```bash
$ chef supermarket search --help
Search indexes allow queries to be made for any type of data that is indexed by the Chef Infra Server, including data bags (and data bag items), environments, nodes, and roles. A defined query syntax is used to support search patterns like exact, wildcard, range, and fuzzy. A search is a full-text query that can be done from several locations, including from within a recipe, by using the search subcommand in chef.

Usage:
  chef supermarket search [flags]

Flags:
  -f, --format string   will be use to search cookbook (default "yaml")
  -h, --help            help for search
  -q, --query string    will be use to search cookbook

Global Flags:
      --config string             config file location (default is $HOME/.chef)
  -m, --supermarket-site string   will be use as cookbook locator (default "https://supermarket.chef.io")
```
### `download` sub-command help
```bash
$ chef supermarket download --help
A cookbook will be downloaded as a tar.gz archive and placed in the current working directory. If a cookbook (or cookbook version) has been deprecated and the --force option is not used, chef will alert the user that the cookbook is deprecated and then will provide the name of the most recent non-deprecated version of that cookbook.

Usage:
  chef supermarket download [flags]

Flags:
  -f, --file string   The filename to write to.
      --force         Force download deprecated version.
  -h, --help          help for download

Global Flags:
      --config string             config file location (default is $HOME/.chef)
  -m, --supermarket-site string   will be use as cookbook locator (default "https://supermarket.chef.io")
```
### `install` sub-command help
```bash
$ chef supermarket install --help
install command will install cookbook that has been downloaded from Chef Supermarket to a local git repository. This action uses the git version control system in conjunction with Chef Supermarket site to install community-contributed cookbooks to the local chef-repo.

Usage:
  chef supermarket install [flags]

Flags:
  -B, --branch string          Default branch to work with. (default "master")
  -o, --cookbook-path string   A colon-separated path to look for cookbooks in.
  -h, --help                   help for install
  -D, --skip-dependencies      Skips automatic dependency installation.
  -b, --use-current-branch     Use the current branch.

Global Flags:
      --config string             config file location (default is $HOME/.chef)
  -m, --supermarket-site string   will be use as cookbook locator (default "https://supermarket.chef.io")

```

## Tasks
### Searching cookbooks
```bash
$ chef supermarket search cookbook name
```

### Download cookbook
```bash
$ chef supermarket download cookbook name
```

### install cookbook
```bash
$ chef supermarket install cookbook name
```