'''
misc utils
'''

from types import ModuleType

import sys

def module_version(module, file=sys.stdout):
    version = hasattr(module, '__version__') and module.__version__
    if version:
        print(f'{module.__name__} ({version})', file=file)

def all_module_versions(namespace, file=sys.stdout):
    for module in filter(lambda m : isinstance(m, ModuleType) and m.__name__ != 'playground_utils', namespace.values()):
        module_version(module, file=file) 
