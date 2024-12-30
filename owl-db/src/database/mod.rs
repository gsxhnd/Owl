use thiserror::Error;
pub mod sqlite;

#[derive(Debug, Error)]
pub enum DatabaseError {
    #[error("Database error: {0}")]
    DbError(String),
}

impl From<sqlx::Error> for DatabaseError {
    fn from(value: sqlx::Error) -> Self {
        DatabaseError::DbError(value.to_string())
    }
}

// pub trait Database {}
