use crate::handler::response::HandlerResponse;
use crate::state::AppState;

use axum::{extract::State, http::StatusCode, Json};
use serde::Serialize;
use tracing::{debug, error};

#[derive(Debug, Clone, Serialize)]
pub struct PingResp {
    version: String,
}

#[utoipa::path(get, path = "/ping", tag = "default", responses())]
pub(crate) async fn ping(state: State<AppState>) -> (StatusCode, Json<HandlerResponse<PingResp>>) {
    debug!("ping request");

    let v = match state.db.version().await {
        Ok(version) => version,
        Err(e) => {
            error!("error: {}", e);
            return HandlerResponse::internal_server_error();
        }
    };

    HandlerResponse::ok(PingResp { version: v })
}
