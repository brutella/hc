#!/bin/bash

printf "GIT ADD...\n"
git add .
printf "OK\n\n"

printf "WRITE THE COMMIT...\n"
read commit
echo "Commit: $commit"
printf "OK\n\n"

printf "GIT COMMIT...\n"
git commit -m "$commit"
printf "OK\n\n"

printf "GIT PUSH...\n"
git push -u origin master -f
printf "OK\n\n"
printf "COMPLETE\n\n\n"
