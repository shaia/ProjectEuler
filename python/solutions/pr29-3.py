import itertools

print (len(set(a ** b for (a, b) in itertools.combinations(range(2, 101) * 2, 2))))