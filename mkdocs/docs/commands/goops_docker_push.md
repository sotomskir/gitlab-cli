## goops docker push

Push docker images to registry

### Synopsis

Push docker images to registry. 
If build context is not one of: master, tags, ^.*-stable$ push will be skipped.
If build is from git tag it will also push image with "stable" tag.
If build is from master branch it will also push image with "latest" tag

```
goops docker push TAG [flags]
```

### Options

```
  -h, --help   help for push
```

### Options inherited from parent commands

```
      --config string   config file (default is $HOME/.goops.yaml)
      --debug           Debug output
      --info            Info output
      --no-color        Disable ANSI color output
      --trace           Trace output
```

### SEE ALSO

* [goops docker](goops_docker.md)	 - Docker integrations

###### Auto generated by spf13/cobra on 12-Apr-2019
