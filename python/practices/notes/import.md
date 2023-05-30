### Overview
Python code is organized into both modules and packages.


### Modules
- **Definition**: An object that serves as an organizational unit of Python code. Modules have a namespace containing arbitrary Python objects. Modules are loaded into Python by the process of importing. (Source)
- In practice, a module usually corresponds to one .py file containing Python code.

```python
import math # math acts as a namespace that keeps all the attributes of the module together
print(math.pi)
print(dir()) # shows what’s in the global namespace
print(dir(math)) # show math namespace
print(math.__dict__["pi"])
print(globals())
```

### Packages
- **Definition**: A Python module which can contain submodules or recursively, subpackages. Technically, a package is a Python module with an __path__ attribute. Note: package is still a module.
- In practice, a package typically corresponds to a file directory containing Python files and other directories. 
- To create a Python package yourself, you create a directory and a file named __init__.py inside it.


### Import Path
When you type import something, Python will look for something a few different places before searching the import path.
In particular, it’ll look in a module cache to see if something has already been imported, and it’ll search among the built-in modules.

```python
import sys
print(sys.path)
```
`sys.path` contains 3 different kinds of locations:
1. The directory of the current script (or the current directory if there’s no script, such as when Python is running interactively)
2. The contents of the PYTHONPATH environment variable
3. Other, installation-dependent directories