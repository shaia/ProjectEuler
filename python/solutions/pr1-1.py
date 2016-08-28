def sum_divied_by(target, divider):
    if divider==0:
        raise ValueError("divider equals to zero")
    q1,r1 = divmod(target,divider)
    q2,r2 = divmod(divider * (q1 * (q1 + 1)), 2)
    return q2

def solve():
    print(sum_divied_by(999,3) + sum_divied_by(999,5) - sum_divied_by(999,15))

if __name__== "__main__":
    print(solve())