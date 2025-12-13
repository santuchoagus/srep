SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
echo "DROP TABLE topics;" | sqlite3 "$SCRIPT_DIR/test.db"
sqlite3 "$SCRIPT_DIR/test.db" < "$SCRIPT_DIR/migrations/001_create_topics.sql"

sqlite3 "$SCRIPT_DIR/test.db" < "$SCRIPT_DIR/migrations/test_seed_topics.sql"
