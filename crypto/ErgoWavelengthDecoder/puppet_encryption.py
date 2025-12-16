from itertools import count
from hashlib import sha256

for puppets in count(start = 1):
  for memories in range(1, puppets + 1):
    for ergo in range(1, puppets + 1):
      if puppets ** 3 + memories ** 3 == 94 * ergo ** 3:
        h = sha256(str(ergo).encode()).hexdigest()
        print(f"FLAG{{{h}}}")
        exit(1)
