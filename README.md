csv-test
========

csv test programs in various languages including evil csv input

Currently implemented in Go and Pyhon.
Conclusion for the current implementations with default settings is that
Python and Go differ in that python parses evil.csv without errors but
introduces a silent data corruption.
Go produces a parse error and let the developer decide what to do with leading
spaces in csv fields.

Go
--

    % go run csv-test.go <evil.csv                    
    01 - 0:"character" 1:"quote"
    02 - 0:"Inigo" 1:"You killed my father\nDarth, I am your father"
    03 - 0:"Buddha" 1:""
    04 - 0:"Dada" 1:"'dodo'"
    05 - 0:"Evil Guy" 1:"\";drop table"
    line 7, column 10: bare " in non-quoted-field
    exit status 1

    % go run csv-test.go -trim_leading_space -lazy_quotes <evil.csv
    01 - 0:"character" 1:"quote"
    02 - 0:"Inigo" 1:"You killed my father\nDarth, I am your father"
    03 - 0:"Buddha" 1:""
    04 - 0:"Dada" 1:"'dodo'"
    05 - 0:"Evil Guy" 1:"\";drop table"
    06 - 0:"Expert" 1:"Trust me, I'm an expert"
    07 - 0:"Balmer" 1:"\"Developers, Developers\""
    08 - 0:"Yoda，Do，do not．do not try"
    09 - 0:"Me" 1:"＂Do not"
    10 - 0:"quote me" 1:"please＂"

Python
------

python2.7 and python3.11 behave the same in my tests with `evil.csv`

    % python2.7 ./csv-test.py <evil.csv
    01 - 0:"character" 1:"quote"
    02 - 0:"Inigo" 1:"You killed my father\nDarth, I am your father"
    03 - 0:"Buddha" 1:""
    04 - 0:"Dada" 1:"'dodo'"
    05 - 0:"Evil Guy" 1:"";drop table"
    06 - 0:"Expert" 1:" "Trust me" 2:" I'm an expert""
    07 - 0:"Balmer" 1:""Developers, Developers""
    08 - 0:"Yoda，Do，do not．do not try"
    09 - 0:"Me" 1:"＂Do not"
    10 - 0:"quote me" 1:" please＂"
    11 -

    % python3 ./csv-test.py --trim_leading_space <evil.csv
    01 -  0:"character" 1:"quote"
    02 -  0:"Inigo" 1:"You killed my father\nDarth, I am your father"
    03 -  0:"Buddha" 1:""
    04 -  0:"Dada" 1:"'dodo'"
    05 -  0:"Evil Guy" 1:"";drop table"
    06 -  0:"Expert" 1:"Trust me, I'm an expert"
    07 -  0:"Balmer" 1:""Developers, Developers""
    08 -  0:"Yoda，Do，do not．do not try"
    09 -  0:"Me" 1:"＂Do not"
    10 -  0:"quote me" 1:"please＂"
    11 -

References/Thx
--------------

http://tools.ietf.org/html/rfc4180

Matthias Wiesmann https://wiesmann.codiferes.net/wordpress/archives/19862 for his evil.csv input
