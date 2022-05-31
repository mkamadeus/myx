#!/bin/bash

# simple script to reset directory to its original form

pushd () {
    command pushd "$@" > /dev/null
}

popd () {
    command popd "$@" > /dev/null
}

projects=("titanic" "mnist" "churn" "catdog")

for d in ${projects[@]}; do
	echo -e "\e[33mremoving files for project [$d]\e[0m"
	pushd "./$d"
		rm -f .dockerignore .gitignore Dockerfile requirements.txt main.py
		rm -rf venv
	popd
done