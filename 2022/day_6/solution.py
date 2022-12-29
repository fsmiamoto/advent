import sys
# assumptions
# - the marker exists in the text

# start = 0
# end = 0
# set = {}


# set = {m,j,q}
# s
#    e
# e - s < 3
# e - s + 1 == len(set)
# e -s == 3 and len(set) == 4
# mjqjpqmgbljsphdztnvjfqwrcgsmlb

# set = {j,q,p, m}
#           s  e
input = sys.stdin.readline()

start = 0
end = 0
chars = set()

# Assuming there is an answer
while True:
    chars.add(input[end])
    while len(chars) < (end-start+1):
        if input[start] != input[end]:
            chars.remove(input[start])
        start += 1

    end += 1

    if len(chars) == 14:
        break

print(end)
