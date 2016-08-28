numbers <- seq(1,999)
answer <- sum(union(numbers[numbers %% 3 == 0],numbers[numbers %% 5 == 0]))
print(answer)
