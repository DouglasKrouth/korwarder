#!/bin/bash
poetry run isort ./korwarder/
poetry run black ./korwarder/
poetry run pyflakes ./korwarder/
