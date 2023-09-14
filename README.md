# go_dev
Go source code repository

# Go development environment

Set-up the Go development environment: 

1. dowload and install Go: https://go.dev/dl/
2. configure paths in .zshrc or .bashrc: add ```export PATH=$PATH:/usr/local/go/bin```
3. configure the GOPATH: ```export GOPATH=$HOME/github_local/go_dev```
4. vim install vim-plug and create a .vimcr file with an example content: 

```
call plug#begin('/.vim/plugged')

Plug 'fatih/vim-go'

call plug#end()

filetype off
filetype plugin indent on

set number
set noswapfile
set noshowmode
set ts=2 sw=2 sts=2 et
set backspace=indent,eol,start

" Map <leader> to comma
let mapleader=","

if has("autocmd")
  autocmd FileType go set ts=2 sw=2 sts=2 noet nolist autowrite
endif
```
 

# libraries:

https://pkg.go.dev/std
