from pprint import pprint


def main():
    r = range(1,100+1)
    a = sum_of_squares(r)
    b = square_of_sums(r)
    pprint(b-a)


def sum_of_squares(r):
    sum = 0
    for n in r:
        sum += n * n
    return sum


def square_of_sums(r):
    sum = 0
    for n in r:
        sum += n
    return sum * sum


if __name__ == "__main__":
    main()