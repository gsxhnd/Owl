use std::path;
use tracing::debug;

use owl_db::database;

#[derive(Debug, Clone)]
pub struct AppState {
    pub db: database::sqlite::Database,
}

impl AppState {
    pub async fn new(data_path: String) -> Self {
        debug!("init app state");
        let db_path = path::Path::new(data_path.as_str())
            .join(".owl")
            .join("owl.db");
        debug!("init db path: {:?}", db_path);

        let db = database::sqlite::Database::new(db_path.as_path()).await;
        AppState { db }
    }
}
