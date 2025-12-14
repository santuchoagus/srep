# Spaced Repetition Bucket List 🪣
A simple CLI tool designed to help you focus on specific topics when you are unsure what to work on. For now, the tool does not use any sophisticated algorithm to select the next topic (yet). There is a future plan for it to be configurable to better support active recall and spaced repetition according to your needs.

## Installation
Uses `sqlite3`, `make` and `golang`
Inside the repository folder execute `make` to build from zero, or `make migrate_build` to rebuild without losing the database which is saved at `$HOME/.local/share/srep/app.db`

## Overview
Help can be viewed at any time with `srep help`. There is currently no logic for tags, so they can be ignored.
- `srep` is the base command. When run with no arguments, it displays the current topic
- `add` Adds a topic with the specified tag `-t <tag>`
- `remove` Removes a topic from the database along with its associated data
- `bucket` Selects a topic from the bucket (planning `-t` support), reusing this command skips it and increments the topic's skipped count
- `complete` Clears the current topic and significantly reduces its skipped count
- `list` Lists all topics

