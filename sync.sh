#!/bin/bash

echo "sync setting.json start"

rm setting.json

mv new.json setting.json

rn new.json

echo "sync setting.json end"