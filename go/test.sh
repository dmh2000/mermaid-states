#!/bin/sh


echo "State1a --> State2a" | go run state-gen.go
echo "State1b --> State2b : description text" | go run state-gen.go
echo "State1c --> State2c:description text" | go run state-gen.go
echo "State1d --> State2d: description text" | go run state-gen.go
echo "State1e --> State2e :description text" | go run state-gen.go
echo "[*] --> State2f" | go run state-gen.go
echo "[*] --> State2g : description text" | go run state-gen.go
echo "[*] --> State2h:description text" | go run state-gen.go
echo "[*] --> State2i: description text" | go run state-gen.go
echo "[*] --> State2j :description text" | go run state-gen.go
echo "State1f --> [*]" | go run state-gen.go
echo "State1g --> [*] : description text" | go run state-gen.go
echo "State1h --> [*]:description text" | go run state-gen.go
echo "State1i --> [*]: description text" | go run state-gen.go
echo "State1j --> [*] :description text" | go run state-gen.go

echo "A --> B: :.,?!@=~" | go run state-gen.go
echo "A --> B::.,?!@=~" | go run state-gen.go
echo "A --> B: : .,?!@=~" | go run state-gen.go
