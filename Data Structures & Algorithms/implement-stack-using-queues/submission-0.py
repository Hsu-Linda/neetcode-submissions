class MyStack:

    def __init__(self):
        self.innerStack = []

    def push(self, x: int) -> None:
        self.innerStack.append(x)

    def pop(self) -> int:
        return self.innerStack.pop()        

    def top(self) -> int:
        return self.innerStack[-1]

    def empty(self) -> bool:
        return 0== len(self.innerStack)


# Your MyStack object will be instantiated and called as such:
# obj = MyStack()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()