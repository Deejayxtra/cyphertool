
file=$(find . | wc -l)
total=$((file * 5))
printf "\\t\\vTotal files * 5: $total\\v\\n"