# 1, 2, 3, 4, 5
# 
import queue

class MyStack:

    def __init__(self):
        self.cur:int = -1
        self.q = queue.Queue()
        self.q2 = queue.Queue()

    def push(self, x: int) -> None:
        self.cur = x
        self.q.put(x)

    def pop(self) -> int:
        item = -1
        while not self.q.empty():
            item = self.q.get()
            if self.q.empty():
                break
            elif self.q.qsize() == 1:
                self.cur = item
            self.q2.put(item)
        self.q, self.q2 = self.q2, self.q
        return item
        
    def top(self) -> int:
        return self.cur

    def empty(self) -> bool:
        return self.q.empty()

# Your MyStack object will be instantiated and called as such:
# obj = MyStack()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()