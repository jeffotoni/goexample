// use std::time::Instant;

// fn main() {
//     let start = Instant::now();
//     // let mut primes = vec![];
//     for i in 2..=100000 {
//         let mut is_prime = true;
//         for j in 3..=(i as f64).sqrt() as i32 {  // Cast i to f64 for sqrt, then back to i32 for comparison
//             if i % j == 0 {
//                 is_prime = false;
//                 break;
//             }
//         }

//         if is_prime {
//             println!("{}", i);
//             // primes.push(i);
//         }
//     }

//     // for prime in primes {
//     //     println!("{}", prime);
//     // }

//     let end = Instant::now();
//     let duration = end - start;  // Corrected: Use subtraction for Instants

//     println!("Time: {:?}", duration);
// }

// use std::time::{Instant, Duration};
use std::time::Instant;

fn main() {
    let start = Instant::now();

    let mut primes = Vec::new();

    for i in 2..=1000000 {
        if i % 2 == 0 {
            if i == 2 {
                primes.push(i);
            } else {
                continue;
            }
        }

        let mut is_prime = true;
        for j in 3..=(i as f64).sqrt() as i32 {
            if i % j == 0 {
                is_prime = false;
                break;
            }
        }

        if is_prime {
            primes.push(i);
        }
    }

    for prime in primes {
        println!("{}", prime);
    }
    let end = Instant::now();
    let duration = end - start;

    println!("Time: {:?}", duration);
}
