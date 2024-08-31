# Falcon

```
        ⠀⠀⠀⠀⠀⠀⣠⣴⣾⣿⣿⣿⣿⣿⣿⣶⣤⣄⠀⠀⠀⠀⠀⠀⠀⠀
        ⠀⠀⠀⠀⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣶⣄⠀⠀⠀⠀⠀
        ⠀⠀⣠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣆⠉⠉⢉⣿⣿⣿⣷⣦⣄⡀⠀
        ⠀⠚⢛⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡄     Welcome to Falcon!
        ⠀⢠⣾⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⠿⠿⠿⠿⣿⡇
        ⢀⣿⡿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡿⠋⠁⠀⠀⠀⠀⠀⠀⠈⠃     Version |  v0.0.1
        ⠸⠁⢀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀     Author  |  Mykola Shaparenko
        ⠀⠀⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡏⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀     GitHub  |  https://github.com/sshaparenko
        ⠀⠀⠀⣿⣿⣿⡿⣿⣿⣿⣿⣿⣿⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀     License |  AGPL 3.0
        ⠀⠀⠀⠹⣿⣿⡇⠈⠻⣿⣿⣿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
        ⠀⠀⠀⠀⠈⠻⡇⠀⠀⠈⠙⠿⣿⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀

    _________    __    __________  _   __
   / ____/   |  / /   / ____/ __ \/ | / /
  / /_  / /| | / /   / /   / / / /  |/ /
 / __/ / ___ |/ /___/ /___/ /_/ / /|  /
/_/   /_/  |_/_____/\____/\____/_/ |_/
```

Falcon is a CLI tool designed specifically for Linux users who prefer using Bash as their shell. The tool provides an intuitive and easy-to-use commands history, making it simpler to search through all your commands.

This repository contains the source code for the Falcon CLI itself. **For those interested in the underlying daemon that powers Falcon's functionality, please visit the Falcond repository**

## Table of Contents

- [Requirements](#requirements)
- [Instalation](#instalation)
- [Usage](#usage)
- [Contributing](#contibuting)
- [License](#license)

## Requirements

- Go version 1.22.2

## Instalation

Clone this repo and run the next commands

```shell
go build .
sudo mv falcon /usr/local/bin
```

## Usage

`falcon run` (in progress) starts the falcon. It will listen and log all the commands from current terminal

`falcon pid` returns PID of current shell

`falcon pid --all` retunrs PID of all shells

`falcon stop` (in progress) stops falcon from listetning current terminal inputs. All data will be accessable with the new run
