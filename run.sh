#!/bin/bash
export DBUSER="client"
export DBPASS="L5LevyxEaFsP3p9TGnPddrY4RcZUHLQ3q8rEwL8K2bddBNzS9Q2TY697EWyX4VTB45HL7bkgpcxzTYnVU2v3HAfJTphMni2AYKVp5BtYUmuGgymK86fz2tYuYs9YSLNh"
export DBHOST="127.0.0.1:5432"
export DBNAME="telemetry"
cd server && 
go run .