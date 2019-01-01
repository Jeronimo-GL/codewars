package codewars.observedPIN;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.Stream;


public class ObservedPin {

	private static Map<Character, List<Character>> altkeys= new HashMap<Character, List<Character>>();
	static {
		altkeys.put('1', Stream.of('2', '4').collect(Collectors.toList()));
		altkeys.put('2', Stream.of('1', '3', '5').collect(Collectors.toList()));
		altkeys.put('3', Stream.of('2', '6').collect(Collectors.toList()));
		altkeys.put('4', Stream.of('1', '5', '7').collect(Collectors.toList()));
		altkeys.put('5', Stream.of('2', '4', '6', '8').collect(Collectors.toList()));
		altkeys.put('6', Stream.of('3', '5', '9').collect(Collectors.toList()));
		altkeys.put('7', Stream.of('4', '8').collect(Collectors.toList()));
		altkeys.put('8', Stream.of('5', '7', '9', '0').collect(Collectors.toList()));
		altkeys.put('9', Stream.of('8', '6').collect(Collectors.toList()));
		altkeys.put('0', Stream.of('8').collect(Collectors.toList()));
	}
	
    public static List<String> getPINs(String observed) {
    	Set<String> combs=new HashSet<String>();
    	combs.add(observed);
    	for(int i=0; i<observed.length(); i++) {
    		combs.addAll(expandList(combs, i));
    	}
		return new ArrayList<String>(combs);
    }
    
        
    private static Set<String> expandList(Set<String> combinations, int position) {
    	Set<String> combs = new HashSet<String>();
    	for (String s: combinations) {
    		List<String> c= expandCombination(s, position);
    		combs.addAll(c);  		
    	}
    	return combs;
    }
    
    private static List<String> expandCombination(String pin, int pos){
    	List<Character> cs=altkeys.get(pin.charAt(pos));
    	List<String> combs=new ArrayList<String>();
    	combs.add(pin);
    	for (Character c: cs) {
    		combs.add(replaceCharAt(pin, c, pos));
    	}
    	return combs;
    }
    
    private static String replaceCharAt(String target, Character newChar, int pos) {
        char[] charArray = target.toCharArray(); 
        charArray[pos] = newChar;    
        String outputString = new String(charArray);
        return outputString;
    }

}
