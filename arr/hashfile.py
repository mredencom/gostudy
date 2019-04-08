import hashlib
# Compare two files for the same


def getHash(f):
    line = f.readline()
    hash = hashlib.md5()
    while (line):
        hash.update(line)
        line = f.readline()
    return hash.hexdigest()


def IsHashEqual(f1, f2):
    str1 = getHash(f1)
    str2 = getHash(f2)
    print(str1,str2)
    return str1 == str2


if __name__ == '__main__':
    f1 = open('D:/v/f2.mp4', 'rb')
    f2 = open('D:/v/f3.mp4', 'rb')

    print(IsHashEqual(f1, f2))
