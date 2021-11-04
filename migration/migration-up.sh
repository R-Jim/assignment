if [ -f ../.env ]
then
  export $(cat ../.env | sed 's/#.*//g' | xargs)
  goose up
fi