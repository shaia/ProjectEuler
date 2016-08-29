def solve():
        sum = 0

        # 531441 is kind of a magic number,  Since 9**5 = 59049, we need at least 5 digits. 5*95 = 295245,
        # so with 5 digits we can make a 6 digit number. 6*95 = 354294
        for number in range(2,355000):
                in_sum = 0
                for digit in str(number):
                        in_sum += int(digit)**5
                if in_sum == number:
                        sum += in_sum

        return sum

print(solve())