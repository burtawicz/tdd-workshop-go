# TDD Workshop

[![CircleCI](https://circleci.com/gh/burtawicz/tdd-workshop-go/tree/main.svg?style=svg&circle-token=b54f4458246611d235827f4c7d6cf77ba94d0e55)](https://circleci.com/gh/burtawicz/tdd-workshop-go/tree/main)

## :question: What is this?
A small API for mutating strings. 

## :bangbang: The problem
Some of our tests started failing, upon investigating the cause we've found that some of our core functions have disappeared.
We have only one hope to save our API - **T**est-**D**riven **D**evelopment!

## :clipboard: Requirements
* `Reverse` must reverse the sequence - `Reverse("ABCD") == "DCBA"`
* `Count` must return the number of elements in the sequence - `Count("ABCD") == 4`
* `Separate` must return the sequence with each element comma separated - `Separate("ABCD") == "A,B,C,D"`
* `EncodeBase64` must encode the sequence to valid Base64 string - `EncodeBase64("ABCD") == "QUJDRA=="`
* `DecodeBase64` must decode a Base64-encoded sequence to a regular string - `DecodeBase64("QUJDRA==") == "ABCD"`
* `EncodeShiftLeft` must shift the sequence left N number of times - `EncodeShiftLeft("ABCD", 2) == "CDAB"`
* `DecodeShiftLeft` must reverse the shifting of sequence left N number of times - `DecodeShiftLeft("CDAB", 2) == "ABCD"`
* `EncodeShiftRight` must shift the sequence right, N number of times - `EncodeShiftRight("ABCD", 3) == "CDBA"`
* `DecodeShiftRight` must reverse the shifting of sequence right N number of times - `DecodeShiftRight("CDBA", 3) == "ABCD"`

## :package: Installation
From the CLI, execute
* `git clone https://github.com/burtawicz/tdd-workshop-go` to clone the repository to your machine
* `cd tdd-workshop-go` to enter the project directory
* `go build` to build the project

## :alembic: Running the tests
From the CLI, execute `go test`
