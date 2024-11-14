#!/bin/sh


echo "State1a --> State2a" | python state-gen.py
echo "State1b --> State2b : description text" | python state-gen.py
echo "State1c --> State2c:description text" | python state-gen.py
echo "State1d --> State2d: description text" | python state-gen.py
echo "State1e --> State2e :description text" | python state-gen.py
echo "[*] --> State2f" | python state-gen.py
echo "[*] --> State2g : description text" | python state-gen.py
echo "[*] --> State2h:description text" | python state-gen.py
echo "[*] --> State2i: description text" | python state-gen.py
echo "[*] --> State2j :description text" | python state-gen.py
echo "State1f --> [*]" | python state-gen.py
echo "State1g --> [*] : description text" | python state-gen.py
echo "State1h --> [*]:description text" | python state-gen.py
echo "State1i --> [*]: description text" | python state-gen.py
echo "State1j --> [*] :description text" | python state-gen.py

echo "A --> B: :.,?!@=~" | python state-gen.py
echo "A --> B::.,?!@=~" | python state-gen.py
echo "A --> B: : .,?!@=~" | python state-gen.py
