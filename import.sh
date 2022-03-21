clear
DB=solvis.db

echo Setup database.
cat solvis.sql | sqlite3 $DB

echo Show schema and number of entries.
sqlite3 $DB ".schema solvis" "select count(*) from solvis" ".exit"

echo Import data.
find out -type f -name '*.txt' | xargs -I% sqlite3 $DB ".import --csv --skip 1 % solvis" ".exit"

echo Show number of entries.
sqlite3 $DB "select count(*) from solvis" ".exit"