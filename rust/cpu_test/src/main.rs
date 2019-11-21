use rayon::prelude::*;
use itertools::{repeat_n, concat};
use time::precise_time_ns;

const NANO_TO_MS: u64 = 1_000_000;
const N_ITERATIONS: i64 = 50; // how many times to repeat PRIMES
const PRIMES: [i64; 6] = [112272535095293,
                          112582705942171,
                          112272535095293,
                          115280095190773,
                          115797848077099,
                          1099726899285419];

fn is_prime(n: i64) -> bool {
    match n {
        n if n < 2      => return false,
        n if n == 2     => return true,
        n if n % 2 == 0 => return false,
        _ => {
            let sqrt_n = (n as f64).sqrt().floor() as i64;
            for i in (3..sqrt_n+1).step_by(2) {
                if n % i == 0 {
                    return false;
                }
            }
            return true; 
        }
    }
}

fn parallel() -> usize {
    // repeat PRIMES N_ITERATION times. result is a Vec<Vec<i64>>. flatten this using concat to a Vec<i64>
    let primes = concat(repeat_n::<_>(PRIMES.to_vec(), N_ITERATIONS as usize).collect::<Vec<Vec<i64>>>());
    primes.par_iter()
          .filter(|i| is_prime(**i))
          .count()
}

fn sequential(quiet: bool) -> usize {
    let mut n_prime = 0;

    for iter in 0..N_ITERATIONS {
        for is_prime in PRIMES.iter().map(|n| is_prime(*n)) {
            if !quiet {
                println!("{} {}", iter, is_prime);
            }

            if is_prime {
                n_prime += 1;
            }
        }
    }

    return n_prime;
}

fn main() {
    let start = precise_time_ns();
    let par = parallel();
    let time_par_ms = (precise_time_ns() - start) / NANO_TO_MS;
    println!("parallel {} {} ms", par, time_par_ms); // 755 ms on 3900x


    let start = precise_time_ns();
    let seq = sequential(true);
    let time_seq_ms = (precise_time_ns() - start) / NANO_TO_MS;
    println!("seq {} {} ms", seq, time_seq_ms); // 8579 ms on 3900x
}


