# Solvis
This project contains some tooling to analyze Solvis' CSV files with Grafana and sqlite. The CSV files are imported into sqlite. Grafana is used to analyze the data.

## Prerequisites
You need the following tools to use this repository:

* Go to run the script
* Sqlite as a backend for Grafana
* Grafana
* Sqlite plugin for Grafana

# Grafana Query
This a an example query for Grafana.

```sql
SELECT 
    time, 
    s01 as "Speicher oben",
    s02 as "Warmwasser",
    s03 as "Speicherreferenz",
    s04 as "Puffer oben",
    s05 as "Solar-VL",
    s06 as "Solar-RL",
    s08 as "Kollektor-Temperatur",
    s09 as "Puffer unten",
    s10 as "Außentemperatur",
    s11 as "Zirkulationstemperatur",
    s12 as "Vorlauf",
    a01 as "Solarpumpe %",
    a02 as "Pumpe WW l/h",
    a03 as "Pumpe HK1 %",
    a04 as "HK2 Überschuss",
    a05 as "Pumpe Zirkulations",
    a08 as "HK1 Mischer auf",
    a09 as "HK2 Mischer zu",
    a12 as "Brenner Stufe 1",
    a13 as "Brenner Stufe 2",
    a16 as "Brenner P %"
from solvis;
```

# Grafana Variables
Add this string as a custom variable to a Grafana dashboard.

```
Speicher oben : s01, Warmwasser : s02, Speicherreference : s03, Heizungsputter oben : s04, Solar VL : s05, Solar RL : s06,Solardruck : s07, Kollektortemperatur : s08, Heizungs Puffer unten : s09, Außentemperatur : s10, Zirkulationstemperatur : s11, Vorlauf : s12
```

# Setup database and import data
Use the command ``go run main.go`` to convert the original files to a modified version that can be imported into sqlite. The original txt files must be in folder `in`. The modified files are in `out`.

Use the script ``import.sh`` to import the generated files into sqlite. 

# Data format
|Position |Name |Description|
|---------|-----|-----------|
|0        |Date |dd.mm.yy   |
|1        |Time |hh:mm:ss   |
|2        |S1   |           |
|3        |S2   |           |
|4        |S3   |           |
|5        |S4   |           |
