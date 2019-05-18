class DynamicConnectivity(object):
# class takes the number of objects n as input, 
# and initializes a data structure with objects numbered from 
# 0 to N-1
    def __init__(self, n):
        self.max=n
        self.groups=[]

# union connects point p with point q
    def union(self, p, q):
        g_p = self.findGroup(p)
        g_q = self.findGroup(q)
        if g_p is not None and g_q is None:
            self.groups[g_p].add(q)
        elif g_p is None and g_q is not None:
            self.groups[g_q].add(p)
        elif g_p is None and g_q is None:
            self.groups.append({p,q})
        elif g_p != g_q:
            for e in self.groups[g_q]:
                self.groups[g_p].add(e)
            self.groups.pop(g_q)


    def findGroup(self, p):
        for g in range(len(self.groups)):
            if p in self.groups[g]:
                return g
        return None
    
# connected checks if point p is connected to point q and returns a boolean
    def connected(self, p, q):
        for g in self.groups:
            if p in g and q in g:
                return True
        return False


def main():
    c1 = DynamicConnectivity(10)
    c1.union(1,3)
    c1.union(3,2)
    c1.union(5,6)
    c1.union(1,5)
    print(c1.connected(1,3))
    print(c1.connected(1,5))

if __name__ == '__main__':
    main()
