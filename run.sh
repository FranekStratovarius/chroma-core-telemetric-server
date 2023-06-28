#!/bin/bash
export DBUSER="root"
export DBHOST="127.0.0.1:3306"
export DBNAME="telemetry"
cd server && 
go run .