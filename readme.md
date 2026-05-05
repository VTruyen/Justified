# Justified

## Goal

This library provides some justfiles examples to make it easier to get started with just.

## Required

Here are the requirements to use this tool:

- `curl` installed on your system.
- `just` installed on your system. (*See this original repository for installation instructions: [Just Repo](https://github.com/casey/just)*)

## Linux/MacOS

### Installation

In order to install this tool, you can run the following command in your terminal:

#### Bash Example

```bash
echo 'jf() { curl -fsSL "https://raw.githubusercontent.com/VTruyen/Justified/refs/heads/master/$1/justfile"; }' >> ~/.bashrc

source ~/.bashrc
```

#### Zsh Example

```bash
echo 'jf() { curl -fsSL "https://raw.githubusercontent.com/VTruyen/Justified/refs/heads/master/$1/justfile"; }' >> ~/.zshrc

source ~/.zshrc
```

### Usage

To use this tool, simply run the following command in your terminal at the root of your project:

```bash
jf <langage> > justfile
```

*Note: Replace `<language>` with the programming language you want to use. For example, if you want to use Python, you would run:*  

```bash
jf python > justfile
```

*Note': C is not supported yet , but it will be added in the future.*

## Windows

*TODO...*
