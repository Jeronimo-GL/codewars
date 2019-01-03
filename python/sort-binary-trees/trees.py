import queue


class Node:
    def __init__(self, L, R, n):
        self.left = L
        self.right = R
        self.value = n


def tree_by_levels(node):
    final = []
    if node is not None:
        final.append(node.value)

        q = queue.Queue()
        if node.left is not None:
            q.put(node.left)
        if node.right is not None:
            q.put(node.right)
        while not q.empty():
            e = q.get()
            final.append(e.value)
            if e.left is not None:
                q.put(e.left)
            if e.right is not None:
                q.put(e.right)

    return final
