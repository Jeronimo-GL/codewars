from functools import reduce


class RomanNumerals:
    parts_values = {'I': 1,
                    'IV': 4,
                    'V': 5,
                    'IX': 9,
                    'X': 10,
                    'IL': 49,
                    'L': 50,
                    'XC': 90,
                    'IC': 99,
                    'C': 100,
                    'CD': 400,
                    'XD': 490,
                    'ID': 499,
                    'D': 500,
                    'CM': 900,
                    'XM': 990,
                    'IM': 999,
                    'M': 1000}
    values = {'I': 1,
              'V': 5,
              'X': 10,
              'L': 50,
              'C': 100,
              'D': 500,
              'M': 1000}

    def from_roman(romStr):
        return reduce(lambda a, b: a + b,
                      map(lambda x: RomanNumerals.parts_values[x],
                          RomanNumerals.parse_roman(romStr)))

    def parse_roman(romStr):
        '''Parses a roman numeral an returns its particles'''
        pos = 0
        parts = []
        while pos < len(romStr):
            new = RomanNumerals.parse_particle(romStr[pos:])
            parts.append(new)
            pos = pos + len(new)
        return parts

    def parse_particle(pstr):
        ''' Returns the next particle in a string '''
        if len(pstr) == 1:
            return pstr
        if RomanNumerals.values[pstr[0]] < RomanNumerals.values[pstr[1]]:
            return pstr[0:2]
        else:
            return pstr[0]

    def to_roman(num):
        romstr = ''
        mils = num // 1000
        romstr += mils*'M'
        num = num % 1000
        cents = num // 100
        if cents == 9:
            romstr += 'CM'
        elif cents >= 5:
            romstr += 'D' + (cents-5)*'C'
        elif cents == 4:
            romstr += 'CD'
        else:
            romstr += cents*'C'
        num = num % 100
        decs = num // 10
        if decs == 9:
            romstr += 'XC'
        elif decs >= 5:
            romstr += 'L' + (decs - 5)*'X'
        elif decs == 4:
            romstr += 'XL'
        else:
            romstr += decs*'X'
        units = num % 10
        if units == 9:
            romstr += 'IX'
        elif units >= 5:
            romstr += 'V' + (units - 5)*'I'
        elif units == 4:
            romstr += 'IV'
        else:
            romstr += units*'I'
        return romstr


if __name__ == "__main__":
    print("Roman numerals")
    num = 8
    print(f"{num} -> {RomanNumerals.to_roman(num)}")
