SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
mkdir "$HOME/.local/share/srep"
touch "$HOME/.local/share/srep/app.db"

sqlite3 "$HOME/.local/share/srep/app.db" < "$SCRIPT_DIR/migrations/001_create_topics.sql"
