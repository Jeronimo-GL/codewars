package codewars

import math.BigInt

object ProperFractions {
    def properFractions(n: Long): Long = {
        if ( n ==1 ) 0 else
        (1L to n).filter(x => n % x != 0).filter(BigInt(_).gcd(n) == 1).length +1
    }
}
