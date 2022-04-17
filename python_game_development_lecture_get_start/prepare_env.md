# prepare environment for opensuse leap 15.x

## network

```bash
opensuse:~ # yast lan
opensuse:~ # yast firewall
```


---

## python

```bash
opensuse:~ # zypper in python3
opensuse:~ # zypper in python39
opensuse:~ $ which python3.6
opensuse:~ $ which python3.9
opensuse:~ $ python3 --version
```


### venv

```bash
opensuse:~ $ python3.9 -m venv <venv>
opensuse:~ $ source py39/bin/activate
(<venv>) opensuse:~ $ which python
(<venv>) opensuse:~ $ python3 --version
(<venv>) opensuse:~ $ deactivate
```


### pip

```bash
opensuse:~ $ which pip
opensuse:~ $ pip help
opensuse:~ $ pip list
opensuse:~ $ pip search <package>
opensuse:~ $ pip install <package>[==<version>]
opensuse:~ $ pip uninstall <package>
```

### ipython

```bash
opensuse:~ $ pip install ipython
opensuse:~ $ ipython
In []: %run <file>.py               # run/load python file
In []: %run -d -b <n> <file>.py     # run python with pdb / debug
ipdb> ?         # help
ipdb> ll        # list code
ipdb> b         # list break point
ipdb> b <n>     # add break point in <n> line
ipdb> cl        # clear all break point
ipdb> cl <m>    # clear <m> break point
ipdb> p <var>   # show var
ipdb> c         # go to next break point
ipdb> q         # q
In []: %%time                       # time execution
...:L = []
...: for n in range(1000):
...:     L.append(n ** 2)
In []: %timeit L = [n ** 2 for n in range(1000)]
In []: %time?
In []: %hist [-n]                   # history
In []: %hist <i> [<j> ...]
In []: %hist <i>-<j>
In []: %hist -f <file>.py
```


### jupyter

```bash
opensuse:~ $ pip install jupyter
opensuse:~ $ jupyter notebook password     # $HOME/.jupyter/jupyter_notebook_config.json
opensuse:~ $ jupyter notebook [--ip=0.0.0.0] [--port=8888]
```


### tk

```bash
opensuse:~ # zypper in python39-tk
opensuse:~ $ python3.9 -m tkinter
```


### pygame

```bash
opensuse:~ $ pip install pygame
opensuse:~ $ python -m pygame.examples.aliens
```


### pylint

```bash
opensuse:~ $ pip install pylint
opensuse:~ $ pylint hello.py
```


### autopep8

```bash
opensuse:~ $ pip install autopep8
opensuse:~ $ autopep8 hello.py
```


---

## vscode

### setting

```json
// settings.json
{
    "files.trimTrailingWhitespace": true,
    "python.linting.enabled": true,
    // linter option: 'bandit', 'flask8', 'mypy', 'prospector',
    //                'pycodestyle', 'pylama', 'pylint'
    // "python.linting.pylintEnabled": true,
    // "python.linting.pylintArgs": ["--disable=C0111"],
    "python.linting.mypyEnabled" : true,
    // formater option: 'autopep8', 'black', 'yapf'
	"python.formatting.provider": "autopep8",
    "[python]":{
        "editor.formatOnType": true,
        "editor.formatOnSave": true,
        "editor.insertSpaces": true,
        "editor.detectIndentation": true,
        "editor.tabSize": 4
    },
    "editor.rulers": [
        {
          "column": 80,
          "color": "#ff9900"
        },
        {
            "column":100,
            "color":"#fbff11"
        },
        {
         "column": 120,
         "color": "#9f0af5"
        },
    ],
}
```
