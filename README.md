<p align="center" >
  <img src="logo.png" alt="logo" width="250"/>
  <h3 align="center">aws-sso (ALPHA)</h3>
  <p align="center">AWS credentials using SSO</p>
</p>

<p align="center" >
  <img alt="Go report card" src="https://goreportcard.com/badge/github.com/nanih98/aws-sso">
  <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/nanih98/aws-sso">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/nanih98/aws-sso">
</p>

<p align="center" >
<a href="https://github.com/nanih98/aws-sso/actions/workflows/releases.yml"><img alt="Pipeline" src="https://github.com/nanih98/aws-sso/actions/workflows/releases.yml/badge.svg"></a>
<a href="https://github.com/nanih98/aws-sso/actions/workflows/lint.yml"><img alt="Pipeline" src="https://github.com/nanih98/aws-sso/actions/workflows/lint.yml/badge.svg"></a>
<a href="https://github.com/nanih98/aws-sso/blob/main/LICENSE"><img alt="LICENSE" src="https://img.shields.io/github/license/nanih98/aws-sso"></a>
</p>

This is a terminal tool to easly log in in AWS using SSO. Build with❤️in Golang

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [About](#about)
  - [Initial requirements](#initial-requirements)
  - [Workflow](#workflow)
- [Installation](#installation)
  - [Using brew](#using-brew)
  - [Using go](#using-go)
  - [Download the binaries](#download-the-binaries)
- [Usage](#usage)
- [Credits](#credits)
- [Contribution](#contribution)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# About
The purpose of this tool is to be able to obtain the credentials accesskey, secretkey, and token of the AWS accounts that your user has access to through SSO.

## Initial requirements
....IN PROGRESS

## Workflow

1. Open your browser and log in using your external identity provider configured with AWS SSO (gmail for example if you use google workspaces as external idp)
2. Install the tool using one of the ways described below

# Installation

## Using brew

```bash
$ brew tap nanih98/aws-sso https://github.com/nanih98/aws-sso
$ brew install aws-sso
```

## Using go

**[Install go](https://go.dev/doc/install)**

```bash
$ go install github.com/nanih98/aws-sso
```
Remember to put your `GOPATH` in your `PATH` variable:

```bash
echo $GOPATH
export PATH=$PATH:/GOPATH/bin # put this in your .bashrc or .zshrc if needed
```

## Download the binaries

[Releases](https://github.com/nanih98/aws-sso/releases)

# Usage

Read [usage](./docs/usage.md) or type:

```bash
$ aws-sso usage
```

# Credits 

- [Github issue](https://github.com/aws/aws-sdk-go-v2/issues/1222)
- [Cobra](https://github.com/spf13/cobra)
- [Charm](https://charm.sh/)
- [Go clipboard](https://github.com/golang-design/clipboard)
- [Promptui](https://github.com/manifoldco/promptui)

# Contribution

Pull requests are welcome! Any code refactoring, improvement, implementation...

# License

[LICENSE](./LICENSE)