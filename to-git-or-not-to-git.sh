#!/bin/bash
infos=$(curl -s https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r '.[] | select(.id == 170) | .name, .powerstats.power, .appearance.gender')  
printf "%s\n" "$infos"