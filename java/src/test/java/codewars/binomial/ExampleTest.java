package codewars.binomial;

import org.junit.Test;
import static org.junit.Assert.assertEquals;

public class ExampleTest { 

	@Test
	public void testBPositive() {        
		assertEquals("1", KataSolution.expand("(x+1)^0"));
		assertEquals("x+1", KataSolution.expand("(x+1)^1"));
		assertEquals("x^2+2x+1", KataSolution.expand("(x+1)^2"));
    }
	
	@Test
	public void testBNegative() {        
		assertEquals("1", KataSolution.expand("(x-1)^0"));
		assertEquals("x-1", KataSolution.expand("(x-1)^1"));
		assertEquals("x^2-2x+1", KataSolution.expand("(x-1)^2"));
	}
	
	@Test
	public void testAPositive() {        
		assertEquals("625m^4+1500m^3+1350m^2+540m+81", KataSolution.expand("(5m+3)^4"));
		assertEquals("8x^3-36x^2+54x-27", KataSolution.expand("(2x-3)^3"));
		assertEquals("1", KataSolution.expand("(7x-7)^0"));
	}
	
	@Test
	public void testANegative() {        
		assertEquals("625m^4-1500m^3+1350m^2-540m+81", KataSolution.expand("(-5m+3)^4"));
		assertEquals("-8k^3-36k^2-54k-27", KataSolution.expand("(-2k-3)^3"));
		assertEquals("1", KataSolution.expand("(-7x-7)^0"));
	}
	
	@Test
	public void testOtherCases() {        
		assertEquals("p^3-3p^2+3p-1", KataSolution.expand("(p-1)^3"));
		assertEquals("64f^6+768f^5+3840f^4+10240f^3+15360f^2+12288f+4096", KataSolution.expand("(2f+4)^6"));
		assertEquals("144t^2-1032t+1849", KataSolution.expand("(-12t+43)^2"));
		assertEquals("r^203", KataSolution.expand("(r+0)^203"));
		assertEquals("x^2+2x+1", KataSolution.expand("(-x-1)^2"));
	}
}
