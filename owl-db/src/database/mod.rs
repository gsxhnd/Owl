use surrealdb::{
    engine::local::{Db, RocksDb},
    Surreal,
};

#[derive(Debug, Clone)]
pub struct Database {
    db: Surreal<Db>,
}

impl Database {
    pub async fn new() -> Self {
        let db = Surreal::new::<RocksDb>("temp.db").await.unwrap();
        db.use_ns("ows").await.unwrap();
        Database { db }
    }
}
