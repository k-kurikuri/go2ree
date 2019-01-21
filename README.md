# go2ree
`go2ee` is tree command implemented by golang

## Installation
This package uses [dep](https://github.com/golang/dep), please install it

```
$ git clone git@github.com:k-kurikuri/go2ree.git && cd go2ree
$ dep ensure
$ go build
```

## Usage
Output the file list of the specified directory
```
$ ./go2ee `specified_dir`
```

If the specified directory is omitted, the file list of the current directory is output
```
$ ./go2ee
```

When the `d` option is specified, only the directory is output
```
$ ./go2ee -d
```

## Unimplemented
- `L` option can not be specified
- Symbolic link is not supported
- Only MacOS supported
