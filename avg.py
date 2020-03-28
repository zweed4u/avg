#!/usr/bin/python3


def avg(*s):
    total = 0
    numbers = len(s)
    for number in s:
        total += number
    return round(total / numbers, 2)


avg(1, 2, 3, 4, 5)
