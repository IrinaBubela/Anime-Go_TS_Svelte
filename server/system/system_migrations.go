package system

func Migrations() error {
    var err error
    
    // Create the favorites table if it doesn't exist
    _, err = Db.Exec(`
        CREATE TABLE IF NOT EXISTS favorites (
            id TEXT PRIMARY KEY NOT NULL,
            anime_id TEXT NOT NULL,
            anime_title TEXT NOT NULL,
            anime_image_url TEXT,
            created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
            updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
        )`)
    if err != nil {
        return err
    }
    
    // Trigger to update 'updated_at' timestamp on UPDATE
    _, err = Db.Exec(`
        CREATE TRIGGER IF NOT EXISTS update_favorites_updated_at
        AFTER UPDATE ON favorites
        FOR EACH ROW
        BEGIN
            UPDATE favorites
            SET updated_at = CURRENT_TIMESTAMP
            WHERE id = OLD.id;
        END;`)
    if err != nil {
        return err
    }
    
    // Add a unique constraint on anime_id to prevent duplicates
    _, err = Db.Exec(`
        CREATE UNIQUE INDEX IF NOT EXISTS idx_unique_anime_id ON favorites (anime_id);
    `)
    if err != nil {
        return err
    }

    return nil
}
