<p align="center" >
  <img src="logo.png" alt="logo" width="250"/>
  <h3 align="center">aws-sso</h3>
  <p align="center">AWS credentials using SSO</p>
</p>

<p align="center" >
  <img alt="Go report card" src="https://goreportcard.com/badge/github.com/nanih98/aws-sso">
  <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/nanih98/aws-sso">
  <img alt="GitHub go.mod Go version" src="https://img.shields.io/github/go-mod/go-version/nanih98/aws-sso">
</p>

[![Pipeline](https://github.com/nanih98/aws-sso/actions/workflows/releases.yml/badge.svg)](https://github.com/nanih98/aws-sso/actions/workflows/releases.yml)
[![Pipeline](https://github.com/nanih98/aws-sso/actions/workflows/lint.yml/badge.svg)](https://github.com/nanih98/aws-sso/actions/workflows/lint.yml)
[![License](https://img.shields.io/github/license/nanih98/aws-sso)](/LICENSE)

This is a terminal tool to easly log in in aws using SSO. Build with❤️in Golang

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [Installation](#installation)
  - [Using go](#using-go)
- [TO DO](#to-do)
- [Credits](#credits)
- [Contributors](#contributors)
- [License](#license)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Installation

## Using go

```bash
$ go install github.com/nanih98/aws-sso
```



# TO DO

- Multithread en las llamadas a los profiles
- El fichero .aws/credentials debe soportar varios profiles de credenciales
- Crear solución nativa de la herramienta para filtrar .aws/credentials y setear el profile que necesitas, reemplazando lo siguiente:

```bash
aws-profile () {
 PROFILE=$(cat ~/.aws/credentials|grep "^\["|sed "s/]$//"|sed "s/^\[//"| fzf)
 export AWS_PROFILE=$PROFILE
}
```
# Credits 

- [Github issue](https://github.com/aws/aws-sdk-go-v2/issues/1222)
- [Cobra](https://github.com/spf13/cobra)

# Contributors

<a href="https://github.com/nanih98/aws-sso/graphs/contributors"><img src="https://opencollective.com/aws-sso/contributors.svg?width=890" /></a>

# License

[LICENSE](./LICENSE)