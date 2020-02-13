[![Build Status](https://travis-ci.org/mpuzanov/parser1c.svg?branch=master)](https://travis-ci.org/mpuzanov/parser1c)
[![Go Report Card](https://goreportcard.com/badge/github.com/mpuzanov/parser1c)](https://goreportcard.com/report/github.com/mpuzanov/parser1c)

# Программа обработки банковского файла из 1С  

## Примеры вызова  

    ./parser1c -format=json example/kl_to_1c.txt
    результат: example/kl_to_1c.json
    ./parser1c -format=csv example/kl_to_1c.txt   
    результат: example/kl_to_1c.csv
    ./parser1c -format=xlsx example/kl_to_1c.txt
    результат: example/kl_to_1c.xlsx
