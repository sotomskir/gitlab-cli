## goops version

Generate semantic version for current HEAD

### Synopsis

Generate semantic version for current HEAD.
Version generation is based on git tags.
If current HEAD is tagged then tag will be used as version.
Else command will lookup for previous tag bump it's minor version, reset patch version and append '-SNAPSHOT'
When there are no tags found version will be '0.1.0-SNAPSHOT'

```
goops version [flags]
```

### Options

```
  -f, --file string   Output file (default "gitlab.env")
  -h, --help          help for version
  -r, --release       Print release version (without -SNAPSHOT)
  -s, --save          Save to file (gitlab.env by default). Override by --file flag
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.goops.yaml)
      --debug           Debug output
      --no-color        Disable ANSI color output
      --trace           Trace output
```

### SEE ALSO

* [goops](goops.md)	 - DevOps toolset written in Go.

###### Auto generated by spf13/cobra on 7-Apr-2019