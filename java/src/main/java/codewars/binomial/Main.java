package codewars.binomial;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Main {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		/*
		String expr="\\((\\-?\\d*)([a-zA-Z])([\\+|\\-]\\d+)?\\)(\\^(\\d+))?";
		
		System.out.println("(3x-23)^2".matches(expr));
		
		Pattern p=Pattern.compile(expr);
		Matcher m=p.matcher("(p)^2");
		
		if(m.find()) {
			System.out.println(m.group(1));
			System.out.println(m.group(2));
			System.out.println(m.group(3));
			System.out.println(m.group(5));
		}
		*/
		
		System.out.println(KataSolution.expand("(7s+17)^10"));
	}

}
