print(abs(3))
print(all([1, 2, 3]))
print(all([1, 2, 3, 0]))
print(any([1, 2, 3, 0]))
print(any([0, ""]))
print(chr(97))
print(dir([1, 2, 3]))
print(divmod(7, 3))
print(eval('1+2'))
print(eval("'hi' + 'a'"))
print(hex(1024))



myList = [lambda a,b:a+b, lambda a,b:a*b]
print(myList)



print(list(zip([1, 2, 3], [4, 5, 6])))
print(list(zip([1, 2, 3], [4, 5, 6], [7, 8, 9])))
print(list(zip("abc", "def")))