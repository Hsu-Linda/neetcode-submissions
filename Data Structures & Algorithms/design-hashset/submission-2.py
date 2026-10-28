class MyHashSet:

    def __init__(self):
        self.store = {}

    def add(self, key: int) -> None:
        self.store[key] = ""

    def remove(self, key: int) -> None:
        self.store.pop(key, None)

    def contains(self, key: int) -> bool:
        return self.store.get(key) is not None


# Your MyHashSet object will be instantiated and called as such:
# obj = MyHashSet()
# obj.add(key)
# obj.remove(key)
# param_3 = obj.contains(key)