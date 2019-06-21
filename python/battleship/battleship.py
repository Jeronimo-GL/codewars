
def validate_battlefield(field):
    # write your magic here
    return False

    
def is_ship_correct(field, pos):
    """ Receives a field and a pos (x,y) and returns its size if it correct
    or 0 it is not. It is NOT correct if it does not have the proper form or if it is 
    in contact with other ship"""
    size = 1
    x,y = pos
    if  (x>0 and y>0 and field[x-1][y-1] == 0) and (x>0 and y<=9 and field[x-1][y+1] == 0) and (x<=9 and y<=9 and field[x+1][y+1] == 0) and (x<=9 and y>=0 and field[x+1][y-1]):
        return size
    else:
        return 0


def main():
    field = [[1, 0, 0, 0, 0, 1, 1, 0, 0, 0],
             [1, 0, 1, 0, 0, 0, 0, 0, 1, 0],
             [1, 0, 1, 0, 1, 1, 1, 0, 1, 0],
             [1, 0, 0, 0, 0, 0, 0, 0, 0, 0],
             [0, 0, 0, 0, 0, 0, 0, 0, 1, 0],
             [0, 0, 0, 0, 1, 1, 1, 0, 0, 0],
             [0, 0, 0, 0, 0, 0, 0, 0, 1, 0],
             [0, 0, 0, 1, 0, 0, 0, 0, 0, 0],
             [0, 0, 0, 0, 0, 0, 0, 1, 0, 0],
             [0, 0, 0, 0, 0, 0, 0, 0, 0, 0]]
    print(is_ship_correct(field, (0,0)))

if __name__ == '__main__':
    main()
