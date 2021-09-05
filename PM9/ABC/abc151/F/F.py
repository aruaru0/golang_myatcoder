from copy import copy
from itertools import combinations
from pprint import pprint
from random import random, sample, seed


class circle:
    def __init__(self, pts):
        assert(0 <= len(pts) and len(pts) <= 3)
        self.__pts = copy(pts)
        if len(pts) == 0:
            self.__cx, self.__cy, self.__r = float('inf'), float('inf'), 0
        elif len(pts) == 1:
            self.__cx, self.__cy, self.__r = *pts[0], 0
        elif len(pts) == 2:
            self.__cx, self.__cy, self.__r = self.__c2()
        else:
            self.__cx, self.__cy, self.__r = self.__c3()

    def __c2(self):
        assert(len(self.__pts) == 2)
        (x0, y0), (x1, y1) = self.__pts
        return ((x0+x1)/2, (y0+y1)/2,
                0.5*((x0-x1)*(x0-x1) + (y0-y1)*(y0-y1))**0.5)

    def __c3(self):
        assert(len(self.__pts) == 3)
        return self.__o3() if self.__obtuse() else self.__a3()

    def __o3(self):
        p, q = max(((xi-xj)**2 + (yi-yj)**2, ((xi, yi), (xj, yj)))
                   for (xi, yi), (xj, yj) in combinations(self.__pts, 2))[1]
        self.__pts = [p, q]
        return self.__c2()

    def __obtuse(self):
        """ Determine whether given tripe forms obtuse, cf.,
                https://math.stackexchange.com/questions/3200316/
        """
        (x0, y0), (x1, y1), (x2, y2) = self.__pts
        a = (x0-x1)**2 + (y0-y1)**2
        b = (x1-x2)**2 + (y1-y2)**2
        c = (x2-x0)**2 + (y2-y0)**2
        return a > b+c or b > c+a or c > a+b

    def __a3(self):
        """ Equation of circle when three points on the circle are given.
            https://www.geeksforgeeks.org/equation-of-circle-when-three-points-on-the-circle-are-given/
        """
        (x0, y0), (x1, y1), (x2, y2) = self.__pts

        x01, x02 = x0 - x1, x0 - x2
        y01, y02 = y0 - y1, y0 - y2

        y20, y10 = y2 - y0, y1 - y0
        x20, x10 = x2 - x0, x1 - x0

        sx02, sy02 = x0*x0 - x2*x2, y0*y0 - y2*y2
        sx10, sy10 = x1*x1 - x0*x0, y1*y1 - y0*y0

        f = (sx02*x01+sy02*x01 + sx10*x02+sy10*x02) / (2*(y20*x01 - y10*x02))
        g = (sx02*y01+sy02*y01 + sx10*y02+sy10*y02) / (2*(x20*y01 - x10*y02))
        c = -x0*x0 - y0*y0 - 2*g*x0 - 2*f*y0

        return (-g, -f, (f*f + g*g - c)**0.5)


    def contains(self, x, y):
        return ((x-self.__cx)**2 + (y-self.__cy)**2)**0.5 <= self.__r + 1e-9

    def pts(self):
        return set(self.__pts)

    @property
    def center(self):
        return (self.__cx, self.__cy)

    @property
    def radius(self):
        return self.__r


def welzl(p, r=set()):
    """ Find the smallest enclosing circle for given a set of points in plane.
        from https://en.wikipedia.org/wiki/Smallest-circle_problem
    
        input:  Finite sets P and R of points in the plane |R| <= 3.
        output: Minimal disk enclosing P with R on the boundary.
    """
    if len(p) == 0 or len(r) == 3:
        return circle(list(r))
    q = sample(p, 1)[0]   # randomly and uniformly
    d = welzl(set(p) - {q}, r)
    if d.contains(*q):
        return d
    return welzl(set(p) - {q}, r | {q})


def msw(pts):
    """ Matousek, Sharir, Welzl's Algorithm for the minimum-circle problem """
    if len(pts) <= 3:
        return circle(pts)
    return _msw(set(pts), set(sample(pts, 3)))


def _msw(p, r):
    """ Find the circle in recursively
        input:  Finite sets P and R of points in the plane |R| = 3.
        output: Minimal disk enclosing P U R.
        from wikipedia.
    """
    if len(p) == 0:
        return circle(list(r))
    q = sample(p, 1)[0]
    d = _msw(p - {q}, r)
    if d.contains(*q):
        return d
    x = (r | {q}) - welzl(r | {q}).pts()
    return _msw((p - {q}) | x, (r | {q}) - x)

N = int(input())
p = []
for _ in range(N) :
    x, y = map(int, input().split())
    p.append((x,y))

print(msw(p).radius)