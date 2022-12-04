#!/bin/env bash

git pull
systemctl restart sign
systemctl status sign
