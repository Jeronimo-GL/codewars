package codewars.differentiate;

import org.junit.Test;
import static org.junit.Assert.assertEquals;

import java.math.BigInteger;

public class DifferentiateTest {

	@Test
	public void TestDefined() {
		assertEquals(BigInteger.valueOf(12), Differntiate.differentiate("12x+2", 3));
		assertEquals(BigInteger.valueOf(5), Differntiate.differentiate("x^2-x", 3));
		assertEquals(BigInteger.valueOf(-20), Differntiate.differentiate("-5x^2+10x+4", 3));
		assertEquals(BigInteger.valueOf(1), Differntiate.differentiate("x", 3));
		assertEquals(BigInteger.valueOf(3), Differntiate.differentiate("3x", 3));
		assertEquals(BigInteger.valueOf(0), Differntiate.differentiate("1", 3));
		assertEquals(BigInteger.valueOf(0), Differntiate.differentiate("1", 1));
		assertEquals(BigInteger.valueOf(0), Differntiate.differentiate("34", -1));
	}
}
