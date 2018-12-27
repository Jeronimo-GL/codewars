package codewars.binomial;

import java.util.regex.Matcher;	
import java.util.regex.Pattern;
import java.math.*;


public class KataSolution {
	public static String expand(String expr) {
		String regex="\\((\\-?\\d*)([a-zA-Z])([\\+|\\-]\\d+)?\\)(\\^(\\d+))?";
		
		Pattern p=Pattern.compile(regex);
		Matcher m=p.matcher(expr);
		String expanded="";
		if(m.find()) {
			int a=parse(m.group(1),1);
			String var=m.group(2);
			int b=parse(m.group(3),1);
			int e=parse(m.group(5), 1);
			if (e==0) return "1";
			for (int i=0; i<=e; i++) {
				double coef=binomial(e,i)*Math.pow(a,(e-i))*Math.pow(b,i);
				if (i>0 && coef>0) expanded +="+";
				if (Math.abs(coef) == 1 && (e-i) == 0) expanded += String.format("%d", (long) coef);
				if (coef != 1 && coef != -1 && coef != 0) expanded += String.format("%d", (long) coef);
				if (coef ==-1 && e-i > 0) expanded +="-";
				if (e-i >=1 && coef != 0) expanded += var;
				if (e-i > 1 && coef != 0) expanded += String.format("^%d", e-i);
			}
		}
		return expanded;
	}
	
	private static long binomial(int n, int k)
    {
        if (k>n-k)
            k=n-k;
 
        long b=1;
        for (int i=1, m=n; i<=k; i++, m--)
            b=b*m/i;
        return b;
    }
	
	private static int parse(String expr, int defaultvalue) {
		if (expr.equals("-")) expr="-1";
		if(expr!=null && !expr.trim().equals("")) {
			return Integer.parseInt(expr);
		}else {
			return defaultvalue;
		}
	}
	
}
