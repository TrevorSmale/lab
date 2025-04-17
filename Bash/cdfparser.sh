#!/bin/bash

# IFS is Internal Field Seprator
# While identified field seperator exists
# Loop over each line and look for example "lname" 
# Then echo a mix of variables and words

while IFS=: read -r lname fname license city status; do
    echo "$fname, $lname has license $license, lives in $city, and is $status."
done < parse1.txt
