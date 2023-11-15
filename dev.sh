#!/bin/sh

mix site.build

cd output && python3 -m http.server 8000
