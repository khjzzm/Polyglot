import sys
from copy import copy

a = ["1","2","3"]
# b = copy(a)
b = a
print(b is a)


print("안녕하세요");

a = 3
print(sys.getrefcount(3))
b = 3
print(sys.getrefcount(3))
c = 3
print(sys.getrefcount(3))
d = 3
print(sys.getrefcount(3))

del(a)
print(sys.getrefcount(3))