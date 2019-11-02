from concurrent import futures
from functools import wraps

import os
import math
import time

# from py docs
PRIMES = [
    112272535095293,
    112582705942171,
    112272535095293,
    115280095190773,
    115797848077099,
    1099726899285419]

def simpletimer(f):
    @wraps(f)
    def wrapper(*args, **kwargs):
        start = time.time()
        result = f(*args, **kwargs)
        overall = time.time() - start
        print(f'{f.__name__} took {overall} seconds')
        return result

    return wrapper

def is_prime(n):
    if n < 2:
        return False
    if n == 2:
        return True
    if n % 2 == 0:
        return False

    sqrt_n = int(math.floor(math.sqrt(n)))
    for i in range(3, sqrt_n + 1, 2):
        if n % i == 0:
            return False
    return True

@simpletimer
def parallel(function, *args, quiet=False, process=True):
    executor_impl = futures.ProcessPoolExecutor if process else futures.ThreadPoolExecutor
    with executor_impl(max_workers=None) as executor:
        print(f'Workers: {executor._max_workers}, Jobs: {len(args)} executor: {executor_impl.__name__}')
        for number, result in zip(args, executor.map(function, args)):
            if not quiet: print(f'{function.__name__}({number}) = {result}')

@simpletimer
def sequential(function, *args, quiet=False):
    for number, result in zip(args, map(function, args)):
        if not quiet: print(f'{function.__name__}({number}) = {result}')

def main():
    parallel(is_prime, *(PRIMES*50), quiet=True, process=True)  # 42s
    parallel(is_prime, *(PRIMES*50), quiet=True, process=False) # 150s; GIL makes this horrible
    sequential(is_prime, *(PRIMES*50), quiet=True)              # 144s! faster than parallel w/ threads

if __name__ == '__main__':
    main()
