'''
misc utils
'''
from pathlib import Path
from pprint import pprint
from types import ModuleType
from tqdm import tqdm

import collections
import sys
import os


IMAGENET_MEAN = [0.485, 0.456, 0.406]
IMAGENET_STD  = [0.229, 0.224, 0.225]
IMAGENET_NORMALIZATION = dict(mean=IMAGENET_MEAN, std=IMAGENET_STD)
INVERSE_IMAGENET_MEAN = [-mean / std for mean, std in zip(IMAGENET_MEAN, IMAGENET_STD)]
INVERSE_IMAGENET_STD  = [1.0 / std for std in IMAGENET_STD]
INVERSE_IMAGENET_NORMALIZATION = dict(mean=INVERSE_IMAGENET_MEAN, std=INVERSE_IMAGENET_STD)

def get_dataset_distribution(ds, classes=None, progress=True):
    if hasattr(ds, 'targets'): # if we already have targets, just use those
        targets = ds.targets
    else:
        indices = range(len(ds))
        if progress:
            indices = tqdm(indices)
        targets = (ds[idx][-1] for idx in indices)

    if classes:
       targets = (classes[target] for target in targets)
    
    return collections.Counter(targets)

def print_and_return(x, file=sys.stdout):
    print(x, file=file)
    return x

def module_version(module, file=sys.stdout):
    version = hasattr(module, '__version__') and module.__version__
    if version:
        print(f'{module.__name__} ({version})', file=file)

def all_module_versions(namespace, file=sys.stdout):
    for module in filter(lambda m : isinstance(m, ModuleType) and m.__name__ != 'playground_utils', namespace.values()):
        module_version(module, file=file) 
