#!/bin/sh


echo "State1a --> State2a" | go run . -v
echo "State1b --> State2b : description text" | go run . -v
echo "State1c --> State2c:description text" | go run . -v
echo "State1d --> State2d: description text" | go run . -v
echo "State1e --> State2e :description text" | go run . -v
echo "[*] --> State2f" | go run . -v
echo "[*] --> State2g : description text" | go run . -v
echo "[*] --> State2h:description text" | go run . -v
echo "[*] --> State2i: description text" | go run . -v
echo "[*] --> State2j :description text" | go run . -v
echo "State1f --> [*]" | go run . -v
echo "State1g --> [*] : description text" | go run . -v
echo "State1h --> [*]:description text" | go run . -v
echo "State1i --> [*]: description text" | go run . -v
echo "State1j --> [*] :description text" | go run . -v
echo "A --> B: :.,?!@=~" | go run . -v
echo "A --> B::.,?!@=~" | go run . -v
echo "A --> B: : .,?!@=~" | go run . -v
echo "the following test should fail"
echo "123" | go run . -v || true
