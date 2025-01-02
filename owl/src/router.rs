use axum::{routing, Router};

use crate::handler::folder as folder_handler;
use crate::handler::ping::ping;
use crate::state::AppState;

pub async fn routes(state: AppState) -> Router {
    let v1_api = Router::new()
        .route("/tag", routing::get(ping))
        .route(
            "/folder",
            routing::get(ping).post(folder_handler::create_folder),
        )
        .route("/file", routing::get(ping));

    Router::new()
        .route("/ping", routing::get(ping))
        .nest("/api/v1", v1_api)
        .with_state(state)
}
