
# chameleons is int[3], desiredColor is int from 0 to 2
def chameleon(chameleons, desiredColor):
    # esta solución no vale. No se pueden combinar dos camaleones del mismo color.
    count=0
    remain=0
    for i in range(3):
        if i != desiredColor:
            if chameleons[i] %2 == 0:
                count += chameleons[i]/2
            else:
                count += (chameleons[i] -1)/2
                remain += 1
    if remain == 2 :
        count += 1
    elif remain == 1 and (chameleons[desiredColor] + count)> 0:
        count += 2
    elif remain == 1:
        return -1
    return count


if __name__ == '__main__':
    chameleon([0,0,0], 1)
