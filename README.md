# dv - A command line tool to remove AWS default VPC(s)

I got an idea of removing default VPCs from [delete-aws-default-vpc](https://github.com/davidobrien1985/delete-aws-default-vpc).

## Installation

- homebrew

```bash
$ brew install knakayama/tap/dv
```

## Usage

```bash
This command enables you to remove default VPC(s) in all AWS regions.
Aside from that, you can remove a VPC in each region.

Usage:
  dv [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  ls          List VPCs in all AWS regions
  rm          Remove a default VPC in an AWS region
  rmrf        Remove default VPCs in all AWS regions

Flags:
  -h, --help   help for dv
```

## TODOs

- [ ] Use goroutine for `ls` command because it's extremely slow.
