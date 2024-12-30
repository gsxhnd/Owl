use crate::database::DatabaseError;
use sqlx::{
    migrate::Migrator,
    sqlite::{SqliteConnectOptions, SqlitePoolOptions},
    ConnectOptions, Pool, Sqlite,
};
use std::path::Path;

#[derive(Debug, Clone)]
pub struct Database {
    pool: Pool<Sqlite>,
}

impl Database {
    pub async fn new(path: &Path) -> Self {
        let mut conn_opt = SqliteConnectOptions::new();
        conn_opt = conn_opt.filename(path);
        conn_opt = conn_opt.log_statements(log::LevelFilter::Info);
        conn_opt = conn_opt.create_if_missing(true);

        let pool_option = SqlitePoolOptions::new();
        let pool = pool_option.connect_with(conn_opt).await.unwrap();

        Database { pool }
    }

    pub async fn version(&self) -> Result<String, DatabaseError> {
        // match self.db.version().await {
        //     Ok(version) => Ok(version.to_string()),
        //     Err(e) => Err(DatabaseError::DbError(e.to_string())),
        // }
        Ok("".to_string())
    }
}
