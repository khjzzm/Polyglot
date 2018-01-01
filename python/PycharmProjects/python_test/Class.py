class Calculator:
    def __init__(self):
        self.result = 0

    def adder(self, num):
        self.result += num
        return self.result

    def sub(self, num):
        self.result -= num
        return self.result

cal1 = Calculator()
cal2 = Calculator()


class FourCal:
    def __init__(self, first, second):
        self.first = first
        self.second = second
    def setdata(self, first, second):
        self.first = first
        self.second = second
    def sum(self):
        result = self.first + self.second
        return result
    def div(self):
        result = self.first / self.second
        return result

a = FourCal(4,5)
print(a.sum())

a1 = FourCal(4,0)
# print(a1.div())


class MoreFourCal(FourCal):
    def pow(self):
        return self.first ** self.second
    def div(self):
        if self.second == 0:
            return 0
        else:
            return self.first / self.second

b = MoreFourCal(0,10)
# print(b.sum())
print(b.pow())