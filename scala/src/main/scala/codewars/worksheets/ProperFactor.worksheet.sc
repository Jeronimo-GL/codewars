import math.BigInt

val k:Long = 12

for {
    i <- 1L to k
    if (k % i ) != 0
    if (BigInt(i).gcd(k) == 1)

} yield(i)

(1L to k).filter(x => k % x != 0).filter(BigInt(_).gcd(k) == 1).length +1
