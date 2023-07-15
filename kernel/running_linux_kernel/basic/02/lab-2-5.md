# developed environment for text mode

## content

- [os](#os)
- [vim](#vim)
- [vim-plugin](#vim-plugin)
  - [vundle](#vundle)
  - [ctag](#ctag)
  - [cscope](#cscope)
  - [tagbar](#tagbar)
  - [nerdtree](#nerdtree)
  - [ale](#ale)
  - [YouCompleteMe](#youcompleteme)
  - [vim-gutentags](#vim-gutentags)
- [linux-kernel](linux-kernel)
- [ref](#ref)

---

## os

```text
debian 12, x86_64
```

---

## vim

```bash
debian:~ # apt instasll vim-nox


debian:~ $ cat << EOF >> ~/.vimrc
set hlsearch
syntax enable
EOF

```

```text
colorscheme -> /usr/share/vim/vim*/colors/lists
synatx      -> /usr/share/vim/vim*/syntax
```

---

## vim-plugin

### vundle

```bash
debian:~ $ git clone https://github.com/VundleVim/Vundle.vim.git ~/.vim/bundle/Vundle.vim
```

```text
" ~/.vimrc
set nocompatible              " be iMproved, required
filetype off                  " required

" set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()

Plugin 'VundleVim/Vundle.vim'
" Plugin 'https://github.com/VundleVim/Vundle.vim.git'

call vundle#end()            " required
filetype plugin indent on    " required

" Brief help
" :PluginList       - lists configured plugins
" :PluginInstall    - installs plugins; append `!` to update or just :PluginUpdate
" :PluginSearch foo - searches for foo; append `!` to refresh local cache
" :PluginClean      - confirms removal of unused plugins; append `!` to auto-approve removal
"
" :h vundle         - more details or wiki for FAQ
```

### ctag

```bash
debian:~ # apt install universal-ctags
```

```text
" ~/.vimrc
if filereadable("tags")
    set tags=tags
endif
```

```bash
debian:~/project $ ctag -R .
debian:~/project $ ls tags
```

```text
ctrl + ]: jump to the definition of the keyword
crtl + t: jump to [count] older entry
```

### cscope

```bash
debian:~ # apt install cscope
```

```text
" ~/.vimrc
if has("cscope")
    set csprg=/usr/bin/cscope
    set csto=1
    set cst
    set nocsverb
    " add any database in current directory
    if filereadable("cscope.out")
        cs add cscope.out
    endif
    set csverb
endif

:set cscopequickfix=s-,c-,d-,i-,t-,e-

" nmap <C-_>s :cs find s <C-R>=expand("<cword>")<CR><CR>
" nmap <C-_>g :cs find g <C-R>=expand("<cword>")<CR><CR>
" nmap <C-_>c :cs find c <C-R>=expand("<cword>")<CR><CR>
" nmap <C-_>t :cs find t <C-R>=expand("<cword>")<CR><CR>
" nmap <C-_>e :cs find e <C-R>=expand("<cword>")<CR><CR>
" nmap <C-_>f :cs find f <C-R>=expand("<cfile>")<CR><CR>
" nmap <C-_>i :cs find i ^<C-R>=expand("<cfile>")<CR>$<CR>
" nmap <C-_>d :cs find d <C-R>=expand("<cword>")<CR><CR>
```

```bash
debian:~/project $ cscope -Rbq
debian:~/project $ ls cscope.*
cscope.po.out cscope.out cscope.in.out
```

```text
:cscope add {file}|{dir}        : add a new cscope database/connection
:cscope find {type} {name}      : query cscope
    0 or s: Find this C symbol
    1 or g: Find this definition
    2 or d: Find functions called by this function
    3 or c: Find functions calling this function
    4 or t: Find this text string
    6 or e: Find this egrep pattern
    7 or f: Find this file
    8 or i: Find files #including this file
    9 or a: Find places where this symbol is assigned a value
:cscope help
:cscope kill                    : Kill a cscope connection
:cscope reset                   : reinit all cscope connection
:cscope show                    : print all connection
```

### tagbar

```text
" ~/.vimrc
call vundle#begin()

...
Plugin 'preservim/tagbar'

call vundle#end()            " required

" nmap <F8> :TagbarToggle<CR>
```

### nerdtree

```text
" ~/.vimrc
call vundle#begin()

...
Plugin 'preservim/nerdtree'

call vundle#end()            " required

" nnoremap <leader>n :NERDTreeFocus<CR>
" nnoremap <C-n> :NERDTree<CR>
" nnoremap <C-t> :NERDTreeToggle<CR>
" nnoremap <C-f> :NERDTreeFind<CR>
```

### ale

```text
" ~/.vimrc
call vundle#begin()

...
Plugin 'dense-analysis/ale'

call vundle#end()            " required

" disable ale
" let g:ale_linters_explicit=1
" let g:ale_echo_msg_error_str = 'E'
" let g:ale_echo_msg_warning_str = 'W'
" let g:ale_echo_msg_format = '[%linter%] %s [%severity%]'
" let g:ale_sign_error = '>>'
" let g:ale_sign_warning = '--'

"nmap <silent> <C-k> <Plug>(ale_previous_wrap)
"nmap <silent> <C-j> <Plug>(ale_next_wrap)
```

```text
:ALEInfo
```

### YouCompleteMe

```bash
debian:~ # apt install build-essential cmake vim-nox python3-dev
debian:~ # apt install mono-complete golang nodejs openjdk-17-jdk openjdk-17-jre npm # optional

debian:~ $ cd ~/.vim/bundle/YouCompleteMe
debian:~/.vim/bundle/YouCompleteMe $ python3 install.py --help
debian:~/.vim/bundle/YouCompleteMe $ python3 install.py
debian:~/.vim/bundle/YouCompleteMe $ python3 install.py --all
  --clang-completer     Enable C-family semantic completion engine through libclang.
  --clangd-completer    Enable C-family semantic completion engine through clangd lsp server.
  --cs-completer        Enable C# semantic completion engine.
  --go-completer        Enable Go semantic completion engine.
  --rust-completer      Enable Rust semantic completion engine.
  --java-completer      Enable Java semantic completion engine.
  --ts-completer        Enable JavaScript and TypeScript semantic completion engine.
  --system-libclang     Use system libclang instead of downloading one from llvm.org.
  --ninja               Use Ninja build system.
  --all                 Enable all supported completers
```

```text
" ~/.vimrc
call vundle#begin()

...
Plugin 'ycm-core/YouCompleteMe'

call vundle#end()            " required

set encoding=utf-8
```

### vim-gutentags

```text
" ~/.vimrc
call vundle#begin()

...
Plugin 'ludovicchabant/vim-gutentags'

call vundle#end()            " required

let g:gutentags_project_root = ['.root', '.svn', '.git', '.hg', '.project', '.gitignore']
```

---

## linux-kernel

```bash
debian:~ # apt install elpa-ggtags

debian:~ $ curl -OL https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-5.15.120.tar.xz
debian:~ $ tar Jxf linux-5.15.120.tar.xz
debian:~ $ cd linux-5.15.120/


debian:~/linux-5.15.120 $ make help
debian:~/linux-5.15.120 $ ls arch/arm64/configs
debian:~/linux-5.15.120 $ make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- defconfig
debian:~/linux-5.15.120 $ make ARCH=arm64 CROSS_COMPILE=aarch64-linux-gnu- tags cscope gtags

```

---

## ref

[Vundle.vim](https://github.com/VundleVim/Vundle.vim)
[tagbar](https://github.com/preservim/tagbar)
[nerdtree](https://github.com/preservim/nerdtree)
[ale](https://github.com/dense-analysis/ale)
[YouCompleteMe](https://github.com/ycm-core/YouCompleteMe)
[vim-gutentags](https://github.com/ludovicchabant/vim-gutentags)
