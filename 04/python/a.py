#!/usr/bin/env python
# coding: -utf8
import sys

if __name__ == '__main__':
    noun = input("Enter a noun: ")
    verb = input("Enter a verb: ")
    adjective = input("Enter an adjective: ")
    adverb = input("Enter an adverb: ")
    print("Do you %s your %s %s %s? That's hilarious!" %
        (verb, adjective, noun, adverb))
