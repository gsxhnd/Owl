use crate::handler::response::HandlerResponse;
use crate::state::AppState;

use axum::{
    extract::{Json, State},
    http::StatusCode,
};
use serde::{Deserialize, Serialize};
use tracing::{error, info};
use validator::Validate;

#[derive(Debug, Clone, Validate, Deserialize, Serialize)]
pub struct CreateFolderReq {
    #[validate(required, length(min = 1))]
    name: Option<String>,
    #[serde(default)]
    pid: u32,
}

pub async fn create_folder(
    state: State<AppState>,
    Json(payload): Json<CreateFolderReq>,
) -> (StatusCode, Json<HandlerResponse<String>>) {
    info!("creat folder get req: {:?}", payload);

    match payload.validate() {
        Ok(_) => (),
        Err(e) => return HandlerResponse::err(e),
    }

    let folder_name = payload.name.unwrap();
    match state.db.create_folder(&folder_name, payload.pid).await {
        Ok(_) => HandlerResponse::ok("".to_string()),
        Err(e) => {
            error!("create folder in db error: {}", e);
            HandlerResponse::internal_server_error()
        }
    }
}
