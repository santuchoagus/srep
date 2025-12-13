SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

rm "$HOME/.local/share/srep/app.db"

mkdir -p "$HOME/.local/share/srep"
touch "$HOME/.local/share/srep/app.db"

sqlite3 "$HOME/.local/share/srep/app.db" < "$SCRIPT_DIR/migrations/001_create_topics.sql"

sqlite3 "$HOME/.local/share/srep/app.db" < "$SCRIPT_DIR/migrations/002_create_global_state.sql"
