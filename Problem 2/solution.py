from pprint import pprint


def main():
    print("Hello World!")
    sums = compute_fibonacci_and_sum(4000000, 1, 2)
    pprint(sums)


def compute_fibonacci_and_sum(max: int, first: int, second: int):
    terms = [first, second]
    previous = second
    term = first + second
    sum = 0
    if first % 2 == 0:
        sum += first
    if second % 2 == 0:
        sum += second
    while term < max:
        terms.append(term)
        if term % 2 == 0:
            sum += term
        term += previous
        previous = terms[len(terms)-1]

    return sum


if __name__ == "__main__":
    main()