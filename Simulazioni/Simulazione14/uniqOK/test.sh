#!/bin/bash

echo === SX oracolo === DX studente ===
if
 diff -y <(./oracolo.sh uno.input) <(./uniq uno.input)
then
 echo "=== NON ci sono differenze ==="
else
 echo "=== ATTENZIONE: ci sono differenze con l'oracolo! ==="
fi
