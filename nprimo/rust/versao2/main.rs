use std::time::Instant;

fn sieve_of_eratosthenes(n: usize) -> Vec<usize> {
    let mut primes = vec![true; n + 1];
    let mut result = Vec::new();

    primes[0] = false;
    primes[1] = false;

    for p in 2..=n {
        if primes[p] {
            result.push(p);
            for i in (p * p..=n).step_by(p) {
                primes[i] = false;
            }
        }
    }

    result
}

fn main() {
    let n = 1000000;
    let start = Instant::now();
    let primes = sieve_of_eratosthenes(n);
    let duration = start.elapsed();

    println!("Prime numbers up to {} are:", n);
    for prime in &primes {
        println!("{}", prime);
    }

    // println!("Prime numbers up to {} are: {:?}", n, primes);
    println!("Time: {:?}", duration);
}

