#!/usr/bin/python

import csv
import sys
import optparse

parser = optparse.OptionParser()
parser.add_option("--trim_leading_space", dest="trim_leading_space", action='store_true',
        default=False, help="Trim leading space on begin of csv records")
(opts, args) = parser.parse_args()

spamreader = csv.reader(sys.stdin, skipinitialspace=opts.trim_leading_space)
lineno = 1
for row in spamreader:
    elemno = 0
    print lineno, "-",
    for elem in row:
        print elemno, elem,
        elemno += 1
    print
    lineno += 1
