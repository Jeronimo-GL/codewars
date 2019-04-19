package codewars.differentiate;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;


public class Differntiate {
	 
	public static class Term{
		public int coef;
		public int exp;
		
		public Term(int coef, int exp) {
			this.coef=coef;
			this.exp=exp;
		}
		
		public String toString() {
			return "(" + this.coef + ", " + this.exp + ")";
		}
	}
	
    public static BigInteger differentiate(String equation, long x) {
    	List<Term> d = derivate(parsePolynomial(equation));
    	//System.out.println(d);
        // Your code here!
        return evaluate(d, x);
    }
    
    
    public static BigInteger evaluate(List<Term> poly, long p) {
    	BigInteger r = BigInteger.ZERO;
    	for (Term t:poly) {
    		BigInteger c = BigInteger.valueOf(t.coef);
    		BigInteger point = BigInteger.valueOf(p);
    		r = r.add(c.multiply(point.pow(t.exp)));
    	}
    	return r;
    }
    
    
    public static List<Term> derivate(List<Term> poly){
    	List<Term> derived = new ArrayList<Term>();
    	for (Term t:poly) {
    		if (t.exp>0) {
    			derived.add(new Term(t.coef*t.exp, t.exp-1));
    		}
    	}
    	return derived;
    }
    
    public static List<Term> parsePolynomial(String poly) {
    	Pattern p= Pattern.compile("[-|\\+]?[\\d]*x?[\\^]?[\\d*]?");
    	Matcher m= p.matcher(poly);
    	List<Term> parts = new ArrayList<Term>();
    	while (m.find()) {
    		if (m.group().length()!=0) {
    			parts.add((parseTerm(m.group())));
    		}
    	}
    	return parts;
    }
    
    public static Term parseTerm(String t) {
    	Pattern p = Pattern.compile( "([-|\\+]?[\\d]*)(x?)[\\^]?([\\d*]?)");
    	Matcher m = p.matcher(t);
    	if (m.find()) {
    		return new Term(parse_coef(m.group(1)), parse_exp(m.group(3), m.group(2)));
    	} else {
    		return null;
    	}
    }
    
    
    public static int parse_exp(String e, String x) {
    	if (e.equals("") && x.equals("x")) {
    		return 1;
    	} else if (e.equals("") && x.equals("")) {
    		return 0;
    	} else {
    		return Integer.parseInt(e);
    	}
    }
    
    
    public static int parse_coef(String c) {
    	if (c.equals("") || c.equals("+")) {
    		return 1;
    	}else if (c.equals("-")) {
    		return -1;
    	} else {
    		return Integer.parseInt(c);
    	}
    }
    
    public static void main(String[] args) {
    	
    	System.out.println(Differntiate.differentiate("12x^2+x", 3));
    }
}
