use crate::database::sqlite::Database;
use crate::database::DatabaseError;

impl Database {
    pub async fn create_folder(&self, name: &str, parent: u32) -> Result<(), DatabaseError> {
        let mut tx = self.pool.begin().await.unwrap();
        if let Err(e) = sqlx::query("insert into folder (name, pid)  values (?,?)")
            .bind(name)
            .bind(parent)
            .execute(&mut *tx)
            .await
        {
            tx.rollback().await.unwrap();
            return Err(DatabaseError::from(e));
        }

        tx.commit().await?;
        Ok(())
    }

    pub async fn exist_folder(&self, name: &str, parent: u32) -> Result<bool, DatabaseError> {
        let e = sqlx::query("select * from folder where name = ? and parent = ?")
            .bind(name)
            .bind(parent)
            .fetch_one(&self.pool)
            .await?;
        Ok(true)
    }
}
