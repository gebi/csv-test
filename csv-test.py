#!/usr/bin/python

import csv
import sys

spamreader = csv.reader(sys.stdin)
lineno = 1
for row in spamreader:
    elemno = 0
    print lineno, "-",
    for elem in row:
        print elemno, elem,
        elemno += 1
    print
    lineno += 1
