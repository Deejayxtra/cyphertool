if [ "$#" -ne 1 ]; then
  echo "Usage: $0 <number>"
  exit 1
fi

number=$1

if [ "$number" -gt 100 ]; then
  number=100
fi

for ((i=1; i<=$number; i++)); do
  echo "$i times I've printed ayodejiolumuyiwa"
done
